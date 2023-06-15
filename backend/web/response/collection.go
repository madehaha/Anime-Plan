package response

import (
	"backend/ent"
	"time"
)

type GetAllCollectionsResp struct {
	Collections []ent.Collection `json:"Collections"`
}
type CommentResp struct {
	Username    string    `json:"Username"`
	SubjectName string    `json:"SubjectName"`
	Time        time.Time `json:"Time"`
	Comment     string    `json:"Comment"`
}
type CollectionGetResp struct {
	UID       uint32 `json:"uid,omitempty"`
	SubID     int    `json:"sub_id,omitempty"`
	Type      uint8  `json:"type,omitempty"`
	IfComment bool   `json:"if_comment,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Score     int8   `json:"score,omitempty"`
	Time      string `json:"time,omitempty"`
}

func GetCollectionResp(collection *ent.Collection) CollectionGetResp {
	return CollectionGetResp{
		UID:       collection.UID,
		SubID:     collection.SubID,
		Type:      collection.Type,
		Time:      collection.Time,
		IfComment: collection.IfComment,
		Comment:   collection.Comment,
		Score:     collection.Score,
	}
}
