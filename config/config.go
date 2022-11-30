package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	ConfEnv string
	Cfg     *Conf
	AppPort = "8080"
)

type Conf struct {
	Server ServerConf
	Log    LogConf
	Redoc  RedocConf
	Cmd    CmdConf
}

// NewConfig 初始化配置信息
func NewConfig() {
	var (
		configPath = fmt.Sprintf("./config/config.%s.yaml", ConfEnv)
	)
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
