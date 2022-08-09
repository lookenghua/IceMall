package v1

import (
	"github.com/gofiber/fiber/v2"
	. "ice-mall/common/dto"
	"ice-mall/service"
	. "ice-mall/util"
)

// CreateProduct 创建商品接口
func CreateProduct(c *fiber.Ctx) error {
	apiUtil := ApiUtil{
		Ctx: c,
	}
	data := GetBody[CreateProductDto](c)
	userInfo := apiUtil.GetUserInfo()
	userId := userInfo.ID
	// 创建商品
	{
		product, err := service.CreateProduct(data, userId)
		if err != nil {
			return apiUtil.Fail(CreateError, err.Error())
		}
		return apiUtil.Success(product)
	}
}
