package collection

import "backend/internal/model"

type AddOrUpdateCollectionReq struct {
	Type     model.CollectionType `json:"type" validate:"required,lte=5,gte=1"`
	Comment  string               `json:"comment"`
	Score    uint8                `json:"score"`
	EpStatus uint8                `json:"ep_status"`
}
