package subject

import (
	"backend/ent"
	"backend/ent/subject"
	"backend/internal/logger"
	"context"
	"errors"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}
func (m *MysqlRepo) GetSubject(ctx context.Context) (ent.Subjects, error) {
	subjects, err := m.client.Subject.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (m *MysqlRepo) GetSubjectByID(ctx context.Context, id int64) (ent.Subject, error) {
	subject, err := m.client.Subject.Query().Where(subject.ID(int(id))).First(ctx)
	if err != nil {
		return *subject, err
	}
	return *subject, nil
}

func (m *MysqlRepo) CreateSubject(ctx context.Context, subject1 *ent.Subject) error {
	if u, _ := m.client.Subject.Query().Where(subject.Name(subject1.Name)).First(ctx); u != nil {
		logger.Error("created same subject")
		err := errors.New("already created same subject")
		return err
	}
	_, err := m.client.Subject.Create().
		SetName(subject1.Name).
		SetNameCn(subject1.NameCn).SetSubjectType(subject1.SubjectType).
		SetDate(subject1.Date).SetSummary(subject1.Summary).
		SetImage(subject1.Image).Save(ctx)
	if err != nil {
		logger.Error("error to create this subject")
		return err
	}
	return nil
}
