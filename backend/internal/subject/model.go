package subject

import (
	"context"
	"errors"

	"backend/ent"
	"backend/ent/subject"
	"backend/internal/logger"
	subjectReq "backend/web/request/subject"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m MysqlRepo) GetSubject(ctx context.Context) (ent.Subjects, error) {
	subjectEntities, err := m.client.Subject.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return subjectEntities, nil
}

func (m MysqlRepo) GetSubjectByID(ctx context.Context, subjectId uint32) (*ent.Subject, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.ID(subjectId)).First(ctx)
	if err != nil {
		return subjectEntity, err
	}
	return subjectEntity, nil
}

func (m MysqlRepo) CreateSubject(ctx context.Context, req subjectReq.CreateSubjectReq) error {
	if u, _ := m.client.Subject.Query().Where(subject.Name(req.Name)).First(ctx); u != nil {
		logger.Error("created same subject")
		err := errors.New("already created same subject")
		return err
	}
	_, err := m.client.Subject.Create().SetImage(req.Image).SetSummary(req.Summary).SetName(req.Name).
		SetNameCn(req.NameCN).SetDate(req.Date).SetEpisodes(req.Episodes).Save(ctx)
	if err != nil {
		logger.Error("error to create this subject")
		return err
	}
	return nil
}
