package service

import (
	"ice-mall/model"
	"ice-mall/schema"
	"ice-mall/util"
	"mime/multipart"
)

// SaveFileToAnnex 保存文件到附件表
func SaveFileToAnnex(annexNum string, file *multipart.FileHeader, fileName string, url string, userId uint) (*schema.Annex, error) {
	var contentType = file.Header["Content-Type"][0]
	var fileType = util.CovertTypeFromFile(contentType)

	var data = schema.Annex{
		AnnexNum:     annexNum,
		FileName:     fileName,
		FileType:     fileType,
		ContentType:  contentType,
		OriginalName: file.Filename,
		Size:         file.Size,
		Url:          url,
		UserId:       userId,
	}
	return model.CreateAnnex(data)
}
