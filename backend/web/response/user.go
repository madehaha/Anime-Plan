package response

import (
	"backend/ent"
)

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserGetResp struct {
	// ID of the ent.
	UID uint32 `json:"uid,omitempty"`
	// Username holds the value of the "username" field.
	Username *string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Gid holds the value of the "gid" field.
	Gid uint8 `json:"gid,omitempty"`
	// RegisterTime holds the value of the "register_time" field.
	RegisterTime string `json:"register_time,omitempty"`
}

func NewUserGetResp(members *ent.Members) UserGetResp {
	return UserGetResp{
		UID:          members.ID,
		Username:     members.Username,
		Email:        members.Email,
		Nickname:     members.Nickname,
		Avatar:       members.Avatar,
		Gid:          members.Gid,
		RegisterTime: members.RegisterTime,
	}
}
