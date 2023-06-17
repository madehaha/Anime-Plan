package subject

type CreateSubjectReq struct {
	// subject
	Image    string `json:"image"`
	Summary  string `json:"summary" validate:"required"`
	Name     string `json:"name" validate:"required"`
	NameCN   string `json:"name_cn" validate:"required"`
	Episodes uint8  `json:"episodes" validate:"required"`
	// subject field
	Year    uint32 `json:"year" validate:"required"`
	Month   uint8  `json:"month" validate:"required,lte=12,gte=1"`
	Date    uint8  `json:"date" validate:"required,lte=31,gte=1"`
	WeekDay uint8  `json:"weekDay" validate:"required,lte=7,gte=1"`
}
