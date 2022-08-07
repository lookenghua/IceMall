package service

import (
	. "ice-mall/common/dto"
	"ice-mall/model"
	"ice-mall/schema"
)

// CreateProduct 创建商品
func CreateProduct(data CreateProductDto, userId uint) (*schema.Product, error) {
	product := schema.Product{
		Title:          data.Title,
		BannerAnnexId:  data.BannerAnnexId,
		OriginalPrice:  data.OriginalPrice,
		DeliveryMethod: *data.DeliveryMethod,
		Carriage:       data.Carriage,
		Content:        data.Content,
		CreatorId:      userId,
	}
	return model.CreateProduct(product)
}
