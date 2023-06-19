package ctrl

import (
	"context"
	"errors"

	"backend/ent"
	"backend/internal/logger"
	"backend/internal/subject"
	subjectField "backend/internal/subject_field"
	subjectReq "backend/web/request/subject"
)

type SubjectCtrl struct {
	subjectRepo      subject.MysqlRepo
	subjectFieldRepo subjectField.MysqlRepo
}

func NewSubjectCtrl(subjectRepo subject.MysqlRepo, subjectFieldRepo subjectField.MysqlRepo) SubjectCtrl {
	return SubjectCtrl{
		subjectRepo:      subjectRepo,
		subjectFieldRepo: subjectFieldRepo,
	}
}

func (sc SubjectCtrl) GetSubject() (ent.Subjects, error) {
	subjects, err := sc.subjectRepo.GetSubject(context.Background())
	if err != nil {
		logger.Error("Failed to Get subjects")
		return nil, err
	}
	return subjects, nil
}

func (sc SubjectCtrl) GetSubjectByID(subjectId uint32) (*ent.Subject, error) {
	subjectEntity, err := sc.subjectRepo.GetSubjectByID(context.Background(), subjectId)
	if err != nil {
		logger.Error("Failed to Get subjects by id ")
		return subjectEntity, err
	}

	return subjectEntity, nil
}

func (sc SubjectCtrl) CreateSubject(req subjectReq.CreateSubjectReq) error {
	ctx := context.Background()
	var subjectId uint32
	var err error

	if subjectEntity, _ := sc.subjectRepo.GetSubjectByName(ctx, req.Name); subjectEntity == nil {
		logger.Error("The subject already exist")
		return errors.New("subject already exist")
	}
	if subjectId, err = sc.subjectRepo.CreateSubject(ctx, req); err != nil {
		logger.Error("Fail to create subjects")
		return err
	}
	if err := sc.subjectFieldRepo.CreateSubjectField(
		ctx, subjectId, req.Year, req.Month, req.Date, req.WeekDay,
	); err != nil {
		logger.Error("Failed to create subject field")
		return err
	}
	return nil
}
