package schema

import (
	"ice-mall/schema/user"
)

type User struct {
	MixinModel
	Username string    `gorm:"not null;comment:用户名"`
	Role     user.Role `gorm:"not null;default:USER;comment:角色"`
	Phone    string    `gorm:"not null;comment:手机号"`
	Avatar   string    `gorm:"not null;comment:头像"`
	Password string    `gorm:"comment:密码"`
	Score    int       `gorm:"default:0;comment:积分数"`
	Sex      user.Sex  `gorm:"not null;default:0;comment:性别"`
	Token    string    `gorm:"comment:用户token"`
}
