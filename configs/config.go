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
		Host       string `yaml:"host"`
		UserName   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		AuthSource string `yaml:"authSource"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	} `yaml:"redis"`
	Operators []string `yaml:"operators"`

	Environment string
	OperatorSet map[string]bool
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
	opt.OperatorSet = make(map[string]bool)
	for _, operator := range opt.Operators {
		opt.OperatorSet[operator] = true
	}

	return &opt, nil
}
