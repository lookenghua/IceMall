package model

import (
	"ice-mall/schema"
)

// CreateProduct 创建商品
func CreateProduct(p schema.Product) (*schema.Product, error) {
	result := client.Create(&p)
	return &p, result.Error
}
