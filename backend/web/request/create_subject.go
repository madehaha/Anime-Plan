package request

import "backend/internal/model"

type CreateSubjectReq struct {
	Name    string
	NameCN  string
	Image   string
	Type    model.SubjectType
	Date    string
	Summary string
}
