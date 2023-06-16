package collection

import (
	"backend/internal/model"
)

type AddOrUpdateCollectionReq struct {
	Type     model.CollectionType `json:"type" validate:"required"`
	Comment  string               `json:"comment"`
	Score    uint8                `json:"score"`
	EpStatus uint8                `json:"ep_status"`
}
