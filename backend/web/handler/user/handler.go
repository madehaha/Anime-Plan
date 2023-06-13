package user

import (
	"backend/internal/ctrl"
)

type Handler struct {
	ctrl ctrl.UserCtrl
}

func NewUserHandler(ctrl ctrl.UserCtrl) Handler {
	return Handler{
		ctrl: ctrl,
	}
}
