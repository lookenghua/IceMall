package schema

import (
	"ice-mall/schema/product"
)

type Product struct {
	MixinModel
	Title          string                 `gorm:"not null;comment:商品名称" json:"title"`
	BannerAnnexId  string                 `gorm:"comment:banner附件id" json:"bannerAnnexId"`
	OriginalPrice  string                 `gorm:"default:0;comment:原价" json:"originalPrice"`
	DeliveryMethod product.DeliveryMethod `gorm:"not null;comment:配送方式" json:"deliveryMethod"`
	Carriage       *int                   `gorm:"default:0;comment:运费" json:"carriage"`
	Content        string                 `gorm:"type:text;comment:详情" json:"content"`
	Flag           *bool                  `gorm:"default:true;comment:是否上架" json:"flag"`
	CreatorId      uint                   `gorm:"not null;comment:创建者ID" json:"creatorId"`
}
