package collection

type UpdateCollectionReq struct {
	Comment        string `json:"comment"`
	CollectionType uint8  `json:"collectionType"`
	Score          int8   `json:"score" `
}
