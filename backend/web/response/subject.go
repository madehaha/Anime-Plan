package response

import "backend/ent"

type GetSubjectWithFieldResp struct {
	ID          uint32           `json:"id,omitempty"`
	Image       string           `json:"image,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	Name        string           `json:"name,omitempty"`
	Date        string           `json:"date,omitempty"`
	NameCn      string           `json:"name_cn,omitempty"`
	OnHold      uint32           `json:"on_hold,omitempty"`
	Wish        uint32           `json:"wish,omitempty"`
	Doing       uint32           `json:"doing,omitempty"`
	SubjectType uint8            `json:"subject_type,omitempty"`
	Dropped     uint32           `json:"dropped,omitempty"`
	Watched     uint32           `json:"watched,omitempty"`
	Field       ent.SubjectField `json:"field"`
}
type GetSubjectResp struct {
	ID          uint32 `json:"id,omitempty"`
	Image       string `json:"image,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Name        string `json:"name,omitempty"`
	Date        string `json:"date,omitempty"`
	NameCn      string `json:"name_cn,omitempty"`
	OnHold      uint32 `json:"on_hold,omitempty"`
	Wish        uint32 `json:"wish,omitempty"`
	Doing       uint32 `json:"doing,omitempty"`
	SubjectType uint8  `json:"subject_type,omitempty"`
	Dropped     uint32 `json:"dropped,omitempty"`
	Watched     uint32 `json:"watched,omitempty"`
}

func SubjectResp(subject *ent.Subject) GetSubjectResp {
	return GetSubjectResp{
		ID:      subject.ID,
		Image:   subject.Image,
		Summary: subject.Summary,
		Name:    subject.NameCn,
		OnHold:  subject.OnHold,
		Doing:   subject.Doing,
		Dropped: subject.Dropped,
		Wish:    subject.Wish,
		Watched: subject.Watched,
		NameCn:  subject.Name,
	}
}
func NewSubjectResp(subject *ent.Subject, field *ent.SubjectField) GetSubjectWithFieldResp {
	return GetSubjectWithFieldResp{
		ID:      subject.ID,
		Image:   subject.Image,
		Summary: subject.Summary,
		Name:    subject.NameCn,
		OnHold:  subject.OnHold,
		Doing:   subject.Doing,
		Dropped: subject.Dropped,
		Wish:    subject.Wish,
		Watched: subject.Watched,
		NameCn:  subject.Name,
		Field:   *field,
	}
}
