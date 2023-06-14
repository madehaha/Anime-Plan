package user

import (
	"backend/internal/ctrl"
	"backend/internal/user"
)

type Handler struct {
	ctrl     ctrl.UserCtrl
	userRepo user.MysqlRepo
}

func NewUserHandler(ctrl ctrl.UserCtrl, userRepo user.MysqlRepo) Handler {
	return Handler{
		ctrl:     ctrl,
		userRepo: userRepo,
	}
}
