package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

func GenerateToken(iat int64, expired int64, key string, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expired
	for k, v := range payload {
		logx.Info(k)
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	logx.Info(key)
	return token.SignedString([]byte(key))
}
