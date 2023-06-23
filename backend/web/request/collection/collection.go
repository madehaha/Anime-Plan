package collection

import (
	"backend/internal/collection"
)

type AddOrUpdateCollectionReq struct {
	Type       collection.CollectionType `json:"type" validate:"required,lte=5,gte=1"`
	HasComment *bool                     `json:"has_comment" validate:"required"`
	Comment    string                    `json:"comment"`
	Score      uint8                     `json:"score"`
	EpStatus   uint8                     `json:"ep_status"`
}

type SelfCollectionsReq struct {
	Type collection.CollectionType `json:"Type" validate:"required"`
}
