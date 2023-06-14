package subject

import (
	"backend/internal/ctrl"
)

type Handler struct {
	ctrl ctrl.SubjectCtrl
}

func NewSubjectHandler(ctrl ctrl.SubjectCtrl) Handler {
	return Handler{
		ctrl: ctrl,
	}
}
