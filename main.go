package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"blog/common"
	"blog/configs"
	"blog/controllers"
	"blog/middlewares"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/zh-five/xdaemon"
)

// 定义传参
type options struct {
	Daemon      bool   `short:"d" long:"daemon" description:"以 Daemon 模式运行"`
	Signal      string `short:"s" long:"signal" default:"start" description:"当按 Daemon 模式运行时的操作命令，可选值：start, stop, restart"`
	PidFile     string `short:"p" long:"pid" default:"logs/service.pid" description:"当按 Daemon 模式运行时，指定存储进程ID的文件路径"`
	Config      string `short:"c" long:"config" description:"系统配置文件路径，默认值：./internal/configs/config.yaml"`
	Environment string `short:"e" long:"environment" default:"development" description:"系统运行环境，可选值：production, development"`
}

func (options *options) daemon() {
	if options.Daemon {
		// 设置进程ID
		pid := 0
		// 从进程文件中读取进程ID
		pf, err := ioutil.ReadFile(options.PidFile)
		if err == nil {
			bytesBuffer := bytes.NewBuffer(pf)
			pid, _ = strconv.Atoi(bytesBuffer.String())
		}

		// 如果进程ID大于0，说明程序已运行
		if pid > 0 {
			// 如果是启动命令，直接提示后退出
			if options.Signal == "start" || options.Signal == "" {
				log.Println("service has start")
				return
			}

			// 根据进程ID获取进程
			poc, _ := os.FindProcess(pid)
			// 发送退出信号
			_ = poc.Signal(os.Kill)
			// 删除进程文件
			_ = os.Remove(options.PidFile)

			// 如果是退出命令，直接提示后退出
			if options.Signal == "stop" {
				log.Println("service has stop")
				return
			}

			// 如果是重启命令，提示重启中
			if options.Signal == "restart" {
				log.Println("service restarting")
			}
		} else {
			if options.Signal == "start" || options.Signal == "" {
				log.Println("service starting")
			}
		}

		// 启动一个子进程后主程序退出
		_, _ = xdaemon.Background("runtime/logs/daemon.log", true)

		file, _ := os.OpenFile(options.PidFile, os.O_CREATE|os.O_WRONLY, 0666)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		// 生成一个新的进程文件并写入进程ID
		_, err = io.WriteString(file, strconv.Itoa(os.Getpid()))

		log.Println(os.Getpid(), "start service daemon")
	}
}

// 配置http服务
func buildServer(host, port string) *http.Server {
	router := gin.Default()

	// 加载中间件
	router.Use(middlewares.Cors)
	router.Use(middlewares.Sign)
	router.Use(middlewares.CatchError)
	router.Use(middlewares.Connect)

	// 注册路由
	controllers.RegisterRoutes(router)

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func main() {
	// 定义传参
	options := &options{}

	// 获取传参
	p := flags.NewParser(options, flags.Default)
	if _, err := p.Parse(); err != nil {
		return
	}

	// 以 Daemon 模式运行
	options.daemon()

	// 获取配置文件路径
	if options.Config == "" {
		var abPath string
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abPath = path.Dir(filename)
		}
		options.Config = path.Join(abPath, "configs/config.yml")
	}

	// 初始化配置
	config, err := configs.InitConfig(options.Config, options.Environment)
	if err != nil {
		log.Panicln(err)
	}

	// 初始化数据库
	common.InitDatabase(&common.DatabaseParams{
		Host:       config.Database.Host,
		UserName:   config.Database.UserName,
		Password:   config.Database.Password,
		Database:   config.Database.Database,
		AuthSource: config.Database.AuthSource,
	})

	db := common.GetDatabase().Connect()

	_ = db.AutoMigrate(&models.Entries{})
	_ = db.AutoMigrate(&models.Tags{})
	_ = db.AutoMigrate(&models.Categories{})
	_ = db.AutoMigrate(&models.Feedbacks{})
	_ = db.AutoMigrate(&models.Users{})
	_ = db.AutoMigrate(&models.Interactions{})

	// 写入日志的文件
	gin.DisableConsoleColor()
	logFile, _ := os.Create("runtime/logs/service.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 配置 http 服务
	srv := buildServer(config.HTTP.Host, config.HTTP.Port)

	go func() {
		// 启动 http 服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Server exiting")
	}
	log.Println(os.Getpid(), "stop service daemon")
}
