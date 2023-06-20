package user

import "mime/multipart"

type UserRegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Info struct {
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty" validate:"required"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty" validate:"required"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty" validate:"required"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty" validate:"required"`
}
type ModifyInfoReq struct {
	Info     Info                  `form:"info" validate:"required"`
	FileData *multipart.FileHeader `form:"image" validate:"required" swaggerignore:"true"`
}
