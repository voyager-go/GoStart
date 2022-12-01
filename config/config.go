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
	Server   ServerConf   // 服务配置
	Log      LogConf      // 日志配置
	Redoc    RedocConf    // 文档配置
	Cmd      CmdConf      // 命令配置
	DataBase DataBaseConf // 数据库配置
}

// NewConfig 初始化配置信息
func NewConfig() {
	var (
		configPath = fmt.Sprintf("./config/env/config.%s.yaml", ConfEnv)
	)
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(file, &Cfg)
	fmt.Println(Cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
