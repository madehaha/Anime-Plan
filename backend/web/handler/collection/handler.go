package collection

import (
	"backend/internal/ctrl"
)

type Handler struct {
	ctrl ctrl.CollectionCtrl
}

func NewCollectionHandler(ctrl ctrl.CollectionCtrl) Handler {
	return Handler{
		ctrl: ctrl,
	}
}
