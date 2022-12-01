package config

type DataBaseConf struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DbName   string `yaml:"dbName"`
}
