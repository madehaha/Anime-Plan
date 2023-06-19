package response

import (
	"backend/ent"

	"backend/internal/collection"
)

type GetAllCollectionsResp struct {
	Collections []ent.Collection `json:"Collections"`
}

type CommentsResp struct {
	Comments []CommentResp `json:"comments"`
}

type CommentResp struct {
	SubjectId uint32 `json:"SubjectId"`
	MemberID  uint32 `json:"MemberID"`
	Time      string `json:"Time"`
	Comment   string `json:"Comment"`
}

type GetCollectionResp struct {
	ID uint32 `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	SubjectId uint32                    `json:"subject_id"`
	Type      collection.CollectionType `json:"type,omitempty"`
	// HasComment holds the value of the "has_comment" field.
	HasComment bool `json:"has_comment,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Score holds the value of the "score" field.
	Score uint8 `json:"score,omitempty"`
	// AddTime holds the value of the "add_time" field.
	AddTime string `json:"add_time,omitempty"`
	// EpStatus holds the value of the "ep_status" field.
	EpStatus uint8 `json:"ep_status,omitempty"`
}

func NewCollectionResp(collection *ent.Collection) GetCollectionResp {
	return GetCollectionResp{
		ID:         collection.ID,
		SubjectId:  collection.Edges.Subject.ID,
		Type:       collection.Type,
		HasComment: collection.HasComment,
		Comment:    collection.Comment,
		Score:      collection.Score,
		AddTime:    collection.AddTime,
		EpStatus:   collection.EpStatus,
	}
}

func NewCommentResp(collection *ent.Collection) CommentResp {
	return CommentResp{
		SubjectId: collection.SubjectID,
		MemberID:  collection.MemberID,
		Comment:   collection.Comment,
		Time:      collection.AddTime,
	}
}
