package dto

import (
	"github.com/gofiber/fiber/v2"
	"ice-mall/schema/product"
	"ice-mall/util"
)

type CreateProductDto struct {
	Title          string                  `json:"title" validate:"required"`
	OriginalPrice  string                  `json:"originalPrice" validate:"required"`
	DeliveryMethod *product.DeliveryMethod `json:"deliveryMethod" validate:"required"`
	Carriage       *int                    `json:"carriage" validate:"required"`
	Content        string                  `json:"content" validate:"required"`
	BannerAnnexId  string                  `json:"bannerAnnexId"`
}

func (t CreateProductDto) Validate() error {
	if *t.DeliveryMethod == product.DeliveryMethodExpress {
		if *t.Carriage == 0 {
			return fiber.NewError(util.ValidateError, "运费不能为0")
		}
	}
	return nil
}
