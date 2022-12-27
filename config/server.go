package config

type ServerConf struct {
	Mode       string
	Port       string
	QqwryPath  string `yaml:"qqwryPath"`
	UploadPath string `yaml:"uploadPath"`
	AssetsPath string `yaml:"assetsPath"`
}
