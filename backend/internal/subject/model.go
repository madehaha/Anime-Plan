package subject

import (
	"backend/ent"
	"backend/internal/logger"
	"context"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m *MysqlRepo) CreateSubject(ctx context.Context, subject *ent.Subject) error {
	_, err := m.client.Subject.Create().
		SetName(subject.Name).
		SetNameCn(subject.NameCn).SetSubjectType(subject.SubjectType).
		SetDate(subject.Date).SetSummary(subject.Summary).
		SetImage(subject.Image).Save(ctx)
	if err != nil {
		logger.Error("error to create this subject")
		return err
	}
	return nil
}
