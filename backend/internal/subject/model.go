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
