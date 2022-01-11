# 博客系统接口

- 基于 gin + gorm + mysql
- 标准 restful api 设计
- 服务以后台进程运行，可主动停止和重启

# 启动参数

- ` -d ` 使用后台服务，默认 false，开启参数将以后台服务形式运行
- ` -s ` 服务命令，默认值 `start`，可选值 `stop` `restart`
- ` -c ` 配置文件地址，默认值 `internal/configs/config.yml`
- ` -e ` 运行环境，默认值 `development`，可选值 `production`

# 操作示例

- 启动服务 `./blog -d`
- 中止服务 `./blog -d -s stop`
- 重启服务 `./blog -d -s restart`