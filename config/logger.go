package config

type LogConf struct {
	Debug    bool
	FileName string `yaml:"fileName"`
	DirPath  string `yaml:"dirPath"`
}
