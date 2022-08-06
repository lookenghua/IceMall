package service

import (
	"errors"
	"gorm.io/gorm"
	. "ice-mall/common/dto"
	"ice-mall/model"
	"ice-mall/schema"
	"ice-mall/schema/user"
)

// CreateUser 创建用户
func CreateUser(data *CreateUserDto) (*schema.User, error) {
	u, err := model.FindByUsername(data.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if u != nil {
		return nil, errors.New("用户已存在")
	}
	// 用户不存在
	userInfo := schema.User{
		Username: data.Username,
		Password: data.Password,
		Role:     user.RoleUSER,
		Avatar:   data.Avatar,
		Phone:    data.Phone,
	}
	newUser, createErr := model.CreateUser(userInfo)
	if createErr != nil {
		return nil, createErr
	}
	return &newUser, nil
}

// FindByUsername 根据用户名查询用户
func FindByUsername(username string) (*schema.User, error) {
	u, err := model.FindByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return u, nil
}

// FindByToken 根据token查询用户
func FindByToken(token string) (*schema.User, error) {
	u, err := model.FindByToken(token)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &u, nil
}
