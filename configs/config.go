package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Option 配置定义
type Option struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"database"`

	Environment string
}

// InitConfig 初始化配置
func InitConfig(file, env string) (*Option, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var options map[string]Option
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return nil, err
	}

	opt := options[env]
	opt.Environment = env

	return &opt, nil
}
