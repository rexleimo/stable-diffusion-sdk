package jwtutils

import (
	"fmt"
	"stable-diffusion-sdk/utils/config"

	"github.com/golang-jwt/jwt/v5"
)

func SignedString(claims jwt.Claims) (string, error) {
	secret := []byte(config.GetConfig().JwtConfig.Secret)
	// 创建令牌
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置头部信息
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "HS256"

	token.Claims = claims
	// 签名生成令牌
	return token.SignedString(secret)
}

func Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名算法")
		}
		secret := []byte(config.GetConfig().JwtConfig.Secret)
		// 获取密钥
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// 获取载荷信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的令牌")
}
