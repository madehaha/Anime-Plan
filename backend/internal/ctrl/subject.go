package ctrl

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/internal/subject"
	subjectReq "backend/web/request/subject"

	"context"
	"errors"
)

type SubjectCtrl struct {
	Repo subject.MysqlRepo
}

func NewSubjectCtrl(repo subject.MysqlRepo) SubjectCtrl {
	return SubjectCtrl{
		Repo: repo,
	}
}

func (sc SubjectCtrl) GetSubject() (ent.Subjects, error) {
	subjects, err := sc.Repo.GetSubject(context.Background())
	if err != nil {
		logger.Error("Failed to Get subjects")
		return nil, err
	}
	return subjects, nil
}

func (sc SubjectCtrl) GetSubjectByID(subjectId uint32) (*ent.Subject, error) {
	subjectEntity, err := sc.Repo.GetSubjectByID(context.Background(), subjectId)
	if err != nil {
		logger.Error("Failed to Get subjects by id ")
		return subjectEntity, err
	}

	return subjectEntity, nil
}

func (sc SubjectCtrl) CreateSubject(req subjectReq.CreateSubjectReq, gid uint8) error {
	if gid == 0 {
		logger.Error("Failed to create subjects")
		err := errors.New("limit of permission")
		return err
	}
	if err := sc.Repo.CreateSubject(context.Background(), req); err != nil {
		logger.Error("fail to create subjects")
		return err
	}
	return nil
}
