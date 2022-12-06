package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Payload struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

var (
	genErr   = "生成令牌失败:"
	emptyErr = "令牌为空"
	parseErr = "令牌解析失败"
	validErr = "令牌验证失败"
)

// Generate 生成Token
func Generate(secret string, expire int64, user interface{}, issuer string) (string, error) {
	data := Payload{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expire,
			Issuer:    issuer,
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	if token, err := j.SignedString([]byte(secret)); err != nil {
		return "", errors.New(genErr + err.Error())
	} else {
		return token, nil
	}
}

// Parse 解析Token
func Parse(jwtToken, secret string) (*Payload, error) {
	if jwtToken == "" {
		return nil, errors.New(emptyErr)
	}
	token, err := jwt.ParseWithClaims(jwtToken, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New(parseErr)
	}
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New(validErr)
	}
}
