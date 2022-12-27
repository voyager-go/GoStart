package config

type EMailConf struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}
