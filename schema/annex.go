package schema

import (
	"ice-mall/schema/annex"
)

type Annex struct {
	MixinModel
	AnnexNum     string         `gorm:"comment:附件id" json:"annexNum"`
	FileType     annex.FileType `gorm:"default:0;comment:文件类型" json:"fileType"`
	FileName     string         `gorm:"comment:文件名称" json:"fileName"`
	OriginalName string         `gorm:"comment:文件原始名称" json:"originalName"`
	ContentType  string         `gorm:"default:'';comment:文件content-type" json:"contentType"`
	Size         int64          `gorm:"default:0;comment:文件大小" json:"size"`
	Url          string         `gorm:"comment:文件下载地址" json:"url"`
	UserId       uint           `gorm:"not null;comment:上传者用户id" json:"userId"`
}
