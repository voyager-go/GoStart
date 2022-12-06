package config

type JwtConf struct {
	TokenExpire string `yaml:"tokenExpire"`
	TokenKey    string `yaml:"tokenKey"`
	TokenIssuer string `yaml:"tokenIssuer"`
	JwtSecret   string `yaml:"jwtSecret"`
}
