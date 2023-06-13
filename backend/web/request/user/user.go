package user

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
