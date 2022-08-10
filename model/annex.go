package model

import "ice-mall/schema"

// CreateAnnex 保存附件
func CreateAnnex(data schema.Annex) (*schema.Annex, error) {
	result := client.Create(&data)
	return &data, result.Error
}
