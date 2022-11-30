package config

type RedocConf struct {
	Title    string `yaml:"title"`
	DocPath  string `yaml:"docPath"`
	SpecPath string `yaml:"specPath"`
	SpecFile string `yaml:"specFile"`
	Desc     string `yaml:"desc"`
}
