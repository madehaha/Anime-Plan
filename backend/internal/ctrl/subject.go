package ctrl

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/internal/subject"
	collection2 "backend/web/request/collection"
	subject2 "backend/web/request/subject"
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
func (sc SubjectCtrl) GetSubjectByID(id int64) (ent.Subject, error) {
	subject, err := sc.Repo.GetSubjectByID(context.Background(), id)
	if err != nil {
		logger.Error("Failed to Get subjects by id ")
		return subject, err
	}
	return subject, nil
}

//func (sc SubjectCtrl) AddCollection(MemberId uint32, SubjectId int, AddType model.SubjectType) error {
//	err := sc.Repo.AddCollection(context.Background(), MemberId, SubjectId, AddType)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (sc SubjectCtrl) UpdateCollection(MemberId uint32, SubjectId int, req collection2.UpdateCollectionReq) error {
	err := sc.Repo.UpdateCollection(context.Background(), MemberId, SubjectId, req)
	if err != nil {
		return err
	}
	return nil
}

func (sc SubjectCtrl) SearchSubject(Name string) (*ent.Subject, error) {
	Subject, err := sc.Repo.SearchSubjection(context.Background(), Name)
	if err != nil {
		return Subject, err
	}
	return Subject, nil
}

func (sc SubjectCtrl) CreateSubject(req subject2.CreateSubjectReq, gid uint8) error {
	if gid == 0 {
		logger.Error("Failed to create subjects")
		err := errors.New("limit of permission")
		return err
	}
	var subject = ent.Subject{
		Image:       req.Image,
		Name:        req.Name,
		NameCn:      req.NameCN,
		SubjectType: req.Type,
		Date:        req.Date,
		Summary:     req.Summary,
	}
	if err := sc.Repo.CreateSubject(context.Background(), &subject); err != nil {
		logger.Error("fail to create subjects")
		return err
	}
	return nil
}
