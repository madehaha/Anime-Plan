package subject

import "backend/internal/model"

type CreateSubjectReq struct {
	Name    string            `json:"name" validate:"required"`
	NameCN  string            `json:"name_cn" validate:"required"`
	Image   string            `json:"image" validate:"required"`
	Type    model.SubjectType `json:"type" validate:"required"`
	Date    string            `json:"date" validate:"required"`
	Summary string            `json:"summary" validate:"required"`
}
