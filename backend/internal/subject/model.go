package subject

type Model struct {
	ID       uint32 `json:"id,omitempty"`
	Image    string `json:"image,omitempty"`
	Summary  string `json:"summary,omitempty"`
	Name     string `json:"name,omitempty"`
	NameCn   string `json:"name_cn,omitempty"`
	OnHold   uint32 `json:"on_hold"`
	Wish     uint32 `json:"wish"`
	Doing    uint32 `json:"doing"`
	Dropped  uint32 `json:"dropped"`
	Watched  uint32 `json:"watched"`
	Episodes uint8  `json:"episodes,omitempty"`
}

type InitialInfo struct {
	// subject
	Image    string `json:"image"`
	Summary  string `json:"summary"`
	Name     string `json:"name"`
	NameCN   string `json:"name_cn"`
	Episodes uint8  `json:"episodes"`
	// subject field
	Year    uint32 `json:"year"`
	Month   uint8  `json:"month"`
	Date    uint8  `json:"date"`
	WeekDay uint8  `json:"weekDay"`
}

type Subject struct {
	Image string `json:"image"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// NameCn holds the value of the "name_cn" field.
	NameCn string `json:"name_cn"`
	// Episodes holds the value of the "episodes" field.
	Episodes uint8 `json:"episodes"`
	// Wish holds the value of the "wish" field.

	// subject field
	Rate1 uint32 `json:"rate_1"`
	// Rate2 holds the value of the "rate_2" field.
	Rate2 uint32 `json:"rate_2"`
	// Rate3 holds the value of the "rate_3" field.
	Rate3 uint32 `json:"rate_3"`
	// Rate4 holds the value of the "rate_4" field.
	Rate4 uint32 `json:"rate_4"`
	// Rate5 holds the value of the "rate_5" field.
	Rate5 uint32 `json:"rate_5"`
	// Rate6 holds the value of the "rate_6" field.
	Rate6 uint32 `json:"rate_6"`
	// Rate7 holds the value of the "rate_7" field.
	Rate7 uint32 `json:"rate_7"`
	// Rate8 holds the value of the "rate_8" field.
	Rate8 uint32 `json:"rate_8"`
	// Rate9 holds the value of the "rate_9" field.
	Rate9 uint32 `json:"rate_9"`
	// Rate10 holds the value of the "rate_10" field.
	Rate10 uint32 `json:"rate_10"`
	// AverageScore holds the value of the "average_score" field.
	AverageScore float64 `json:"average_score"`
	// Year holds the value of the "year" field.
	Year uint32 `json:"year"`
	// Month holds the value of the "month" field.
	Month uint8 `json:"month"`
	// Date holds the value of the "date" field.
	Date uint8 `json:"date"`
	// Weekday holds the value of the "weekday" field.
	Weekday uint8 `json:"weekday"`
}
