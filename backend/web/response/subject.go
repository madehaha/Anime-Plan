package response

import "backend/ent"

type SubjectGetResp struct {
	ID          int    `json:"id,omitempty"`
	Image       string `json:"image,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Name        string `json:"name,omitempty"`
	Date        string `json:"date,omitempty"`
	NameCn      string `json:"name_cn,omitempty"`
	OnHold      uint32 `json:"on_hold,omitempty"`
	Wish        uint32 `json:"wish,omitempty"`
	Doing       uint32 `json:"doing,omitempty"`
	SubjectType uint8  `json:"subject_type,omitempty"`
	Collect     uint32 `json:"collect,omitempty"`
	Drop        uint32 `json:"drop,omitempty"`
	Watched     uint32 `json:"watched,omitempty"`
}

func NewSubjectResp(sub *ent.Subject) SubjectGetResp {
	return SubjectGetResp{
		ID:          sub.ID,
		Image:       sub.Image,
		Summary:     sub.Summary,
		Name:        sub.NameCn,
		Date:        sub.Date,
		OnHold:      sub.OnHold,
		Doing:       sub.Doing,
		Drop:        sub.Drop,
		Collect:     sub.Collect,
		Wish:        sub.Wish,
		Watched:     sub.Watched,
		SubjectType: sub.SubjectType,
		NameCn:      sub.Name,
	}
}
