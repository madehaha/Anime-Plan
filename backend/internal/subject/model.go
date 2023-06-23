package subject

import "backend/ent"

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
type Middle struct {
	Subject *ent.Subject
	Field   *ent.SubjectField
}
