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
	SubjectId uint32                    `json:"SubjectId"`
	MemberID  uint32                    `json:"MemberID"`
	Time      string                    `json:"Time"`
	Comment   string                    `json:"Comment"`
	Type      collection.CollectionType `json:"Type"`
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

func NewCollectionResp(Collection *ent.Collection) GetCollectionResp {
	return GetCollectionResp{
		ID:         Collection.ID,
		SubjectId:  Collection.Edges.Subject.ID,
		Type:       Collection.Type,
		HasComment: Collection.HasComment,
		Comment:    Collection.Comment,
		Score:      Collection.Score,
		AddTime:    Collection.AddTime,
		EpStatus:   Collection.EpStatus,
	}
}

func NewCommentResp(Collection *ent.Collection) CommentResp {
	return CommentResp{
		SubjectId: Collection.SubjectID,
		MemberID:  Collection.MemberID,
		Comment:   Collection.Comment,
		Time:      Collection.AddTime,
		Type:      Collection.Type,
	}
}
