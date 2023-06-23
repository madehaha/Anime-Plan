package response

import (
	"backend/ent"
	"backend/internal/subject"
	subjectField "backend/internal/subject_field"
)

type GetSubjectWithFieldResp struct {
	Info  subject.Model
	Field subjectField.Model `json:"field"`
}

type GetSubjectResp struct {
	Info subject.Model
}
type SubjectWithField struct {
	Subject *ent.Subject
	Field   *ent.SubjectField
}

func SubjectResp(Subject *ent.Subject) GetSubjectResp {
	return GetSubjectResp{
		Info: GetInfo(Subject),
	}
}
func NewSubjectResp(Field SubjectWithField) GetSubjectWithFieldResp {
	return GetSubjectWithFieldResp{
		Info:  GetInfo(Field.Subject),
		Field: GetField(Field.Field),
	}
}

func GetInfo(Subject *ent.Subject) subject.Model {
	return subject.Model{
		ID:       Subject.ID,
		Image:    Subject.Image,
		Summary:  Subject.Summary,
		Name:     Subject.Name,
		NameCn:   Subject.NameCn,
		Episodes: Subject.Episodes,
		Wish:     Subject.Wish,
		OnHold:   Subject.OnHold,
		Doing:    Subject.Doing,
		Watched:  Subject.Watched,
		Dropped:  Subject.Dropped,
	}
}
func GetField(Field *ent.SubjectField) subjectField.Model {
	return subjectField.Model{
		ID:           Field.ID,
		Rank:         Field.Rank,
		Rate1:        Field.Rate1,
		Rate2:        Field.Rate2,
		Rate3:        Field.Rate3,
		Rate4:        Field.Rate4,
		Rate5:        Field.Rate5,
		Rate6:        Field.Rate6,
		Rate7:        Field.Rate7,
		Rate8:        Field.Rate8,
		Rate9:        Field.Rate9,
		Rate10:       Field.Rate10,
		AverageScore: Field.AverageScore,
		Year:         Field.Year,
		Month:        Field.Month,
		Date:         Field.Date,
		Weekday:      Field.Weekday,
	}
}
