package common

import (
	"gintest/utils"
)

// 项目通用配置
type CommonConfig struct {
	Env   string        `yaml:"env"`
	Port  string        `yaml:"port"`
	Log   *utils.Logger `yaml:"log"`
	Mysql *utils.Mysql  `yaml:"mysql"`
}

var CConfig CommonConfig
