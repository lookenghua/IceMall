package v1

import (
	"github.com/gofiber/fiber/v2"
	. "ice-mall/common/dto"
	"ice-mall/service"
	. "ice-mall/util"
)

func CreateProduct(c *fiber.Ctx) error {
	apiUtil := ApiUtil{
		Ctx: c,
	}
	data := new(CreateProductDto)
	// 校验
	{
		err := apiUtil.ValidateBody(data)
		if err != nil {
			return apiUtil.Fail(ValidateError, err.Error())
		}
		err1 := data.Validate()
		if err1 != nil {
			return apiUtil.Fail(ValidateError, err.Error())
		}
	}
	userInfo := apiUtil.GetUserInfo()
	userId := userInfo.ID
	// 创建商品
	{
		product, err := service.CreateProduct(*data, userId)
		if err != nil {
			return apiUtil.Fail(CreateError, err.Error())
		}
		return apiUtil.Success(product)
	}
}
