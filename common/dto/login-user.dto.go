package dto

type LoginUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
