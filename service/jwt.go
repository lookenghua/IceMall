package service

import (
	"github.com/golang-jwt/jwt/v4"
	"ice-mall/config"
	"ice-mall/model"
	"time"
)

// CreateToken 创建token
func CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// SaveToken 更新token
func SaveToken(id uint, token string) error {
	_, err := model.UpdateToken(id, token)
	if err != nil {
		return err
	}
	return nil
}
