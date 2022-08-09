package v1

import (
	"github.com/gofiber/fiber/v2"
	. "ice-mall/common/dto"
	"ice-mall/schema/user"
	"ice-mall/service"
	. "ice-mall/util"
)

// LoginAdmin 管理员登录
func LoginAdmin(c *fiber.Ctx) error {
	apiUtil := ApiUtil{
		Ctx: c,
	}
	data := GetBody[LoginAdminDto](c)
	// 查询用户
	findUser, findErr := service.FindByUsername(data.Username)
	if findErr != nil {
		return apiUtil.Fail(DataFoundError, findErr.Error())
	}
	// 用户不存在
	if findUser == nil {
		return apiUtil.Fail(DataNotFoundError, "当前用户不存在")
	}
	// 判断是不是管理员角色
	if findUser.Role != user.RoleADMIN {
		return apiUtil.Fail(AuthorityNotRightError, "当前管理员不存在")
	}
	// 判断密码是否正确
	if findUser.Password != data.Password {
		return apiUtil.Fail(PasswordNotMatchedError, "密码错误")
	}
	// 创建token
	token, createErr := service.CreateToken(findUser.ID)
	if createErr != nil {
		return apiUtil.Fail(CreateTokenError, findErr.Error())
	}
	// 存入token
	saveErr := service.SaveToken(findUser.ID, token)
	if saveErr != nil {
		return apiUtil.Fail(SaveDataError, findErr.Error())
	}
	return apiUtil.Success(token)
}
