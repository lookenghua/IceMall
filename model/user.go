package model

import "ice-mall/schema"

// CreateUser 创建用户
func CreateUser(u schema.User) (schema.User, error) {
	result := client.Create(&u)
	return u, result.Error
}

// FindByUsername 根据用户名查询
func FindByUsername(username string) (*schema.User, error) {
	user := schema.User{}
	result := client.Where("username = ?", username).First(&user)
	return &user, result.Error
}

// UpdateToken 更新用户token
func UpdateToken(id uint, t string) (schema.User, error) {
	user := schema.User{}
	user.ID = id
	result := client.Model(&user).Update("token", t)
	return user, result.Error
}

// FindByToken 根据token查询用户
func FindByToken(token string) (schema.User, error) {
	user := schema.User{}
	client.Where("token = ?", token).First(&user)
	return user, nil
}
