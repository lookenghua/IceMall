package dto

type UploadFileDto struct {
	AnnexNum string `json:"annexNum" form:"annexNum" validate:""`
}
