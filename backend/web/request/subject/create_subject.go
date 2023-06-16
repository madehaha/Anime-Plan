package subject

type CreateSubjectReq struct {
	Image    string `json:"image"`
	Summary  string `json:"summary"`
	Name     string `json:"name" validate:"required"`
	NameCN   string `json:"name_cn" validate:"required"`
	Date     string `json:"date" validate:"required"` // release date
	Episodes uint8  `json:"episodes" validate:"required"`
}
