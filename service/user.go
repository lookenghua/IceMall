package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	. "ice-mall/common/dto"
	"ice-mall/config"
	"ice-mall/ent"
	"ice-mall/ent/user"
	"ice-mall/model"
	"time"
)

// CreateUser 创建用户
func CreateUser(data *CreateUserDto) (*ent.User, error) {
	u, err := model.FindByUsername(data.Username)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if u != nil {
		return nil, errors.New("用户已存在")
	}
	// 用户不存在
	userInfo := ent.User{
		Username: data.Username,
		Password: &data.Password,
		Role:     user.RoleUSER,
		Avatar:   data.Avatar,
		Phone:    &data.Phone,
	}
	newUser, createErr := model.CreateUser(userInfo)
	if createErr != nil {
		return nil, createErr
	}
	return newUser, nil
}

// FindByUsername 根据用户名查询用户
func FindByUsername(username string) (*ent.User, error) {
	u, err := model.FindByUsername(username)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return u, nil
}

// CreateToken 创建token
func CreateToken(id int) (string, error) {
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
func SaveToken(id int, token string) error {
	_, err := model.UpdateToken(id, token)
	if err != nil {
		return err
	}
	return nil
}

// FindByToken 根据token查询用户
func FindByToken(token string) (*ent.User, error) {
	u, err := model.FindByToken(token)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return u, nil
}
