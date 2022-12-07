package config

type JwtConf struct {
	TokenExpire int64  `yaml:"tokenExpire"`
	TokenKey    string `yaml:"tokenKey"`
	TokenIssuer string `yaml:"tokenIssuer"`
	JwtSecret   string `yaml:"jwtSecret"`
}
