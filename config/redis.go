package config

type RedisConf struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	DbNum       int    `yaml:"dbNum"`
	LoginPrefix string `yaml:"loginPrefix"`
}
