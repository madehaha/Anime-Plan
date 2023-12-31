package ctrl

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"backend/ent"
	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/subject"
	subjectField "backend/internal/subject_field"
	subjectReq "backend/web/request/subject"
)

type SubjectCtrl struct {
	subjectRepo      subject.MysqlRepo
	subjectFieldRepo subjectField.MysqlRepo
	cfg              config.AppConfig
}

func NewSubjectCtrl(
	subjectRepo subject.MysqlRepo, subjectFieldRepo subjectField.MysqlRepo, cfg config.AppConfig,
) SubjectCtrl {
	return SubjectCtrl{
		subjectRepo:      subjectRepo,
		subjectFieldRepo: subjectFieldRepo,
		cfg:              cfg,
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

func (sc SubjectCtrl) GetSubjectByID(subjectId uint32) (*ent.Subject, *ent.SubjectField, error) {
	subjectEntity, Field, err := sc.subjectRepo.GetSubjectByID(context.Background(), subjectId)
	if err != nil {
		logger.Error("Failed to Get subjects by id ")
		return subjectEntity, Field, err
	}

	return subjectEntity, Field, nil
}
func (sc SubjectCtrl) GetSubjectByName(name string) (*ent.Subject, *ent.SubjectField, error) {
	subjectEntity1, field1, err1 := sc.subjectRepo.GetSubjectByNameCN(context.Background(), name)
	subjectEntity2, field2, err2 := sc.subjectRepo.GetSubjectByName(context.Background(), name)
	if (err1 != nil) && (err2 != nil) {
		logger.Error("Failed to Get subjects by id ")
		return nil, nil, err1
	}
	if err1 != nil {
		return subjectEntity2, field2, err2
	}
	if err2 != nil {
		return subjectEntity1, field1, err1
	}
	if subjectEntity1.ID == subjectEntity2.ID {
		return subjectEntity1, field1, err1
	}

	return nil, nil, errors.New("compare error")
}
func (sc SubjectCtrl) Rankings() ([]subject.Middle, error) {
	res, err := sc.subjectRepo.Rankings(context.Background())
	if err != nil {
		return nil, err
	}
	return res, err
}
func (sc SubjectCtrl) CreateSubject(req subjectReq.CreateSubjectReq) error {
	ctx := context.Background()
	var subjectId uint32
	var err error

	if subjectEntity, _, _ := sc.subjectRepo.GetSubjectByName(ctx, req.Name); subjectEntity != nil {
		logger.Error("The subject already exist")
		return errors.New("subject already exist")
	}
	info := subjectReq.NewInitialInfo(req)
	if subjectId, err = sc.subjectRepo.CreateSubject(ctx, info); err != nil {
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

func (sc SubjectCtrl) BoardCast(Day uint8) (subjects ent.Subjects, err error) {
	subjects, err = sc.subjectRepo.BoardCast(context.Background(), Day)
	return
}
func (sc SubjectCtrl) CreateSubjectWithSave(uid uint32, req subjectReq.CreateSubjectWithSaveReq) error {
	ctx := context.Background()
	var subjectId uint32
	var err error
	if subjectEntity, _, _ := sc.subjectRepo.GetSubjectByName(ctx, req.CreateSubject.Name); subjectEntity != nil {
		logger.Error("The subject already exist")
		return errors.New("subject already exist")
	}
	fileHeader := req.FileData
	filename := strconv.FormatUint(uint64(uid), 10) + "_" + time.Now().Format(
		"200601021504",
	) + "_" + filepath.Ext(
		fileHeader.Filename,
	)
	err = sc.saveFile(fileHeader, filename)
	if err != nil {
		logger.Error("Failed to save file")
		return err
	}
	imageUrl := sc.cfg.WebDomain + "/" + sc.cfg.StaticDirectory + "/" + filename
	logger.Info(imageUrl)
	req.CreateSubject.Image = imageUrl
	info := subjectReq.NewInitialInfo(req.CreateSubject)
	if subjectId, err = sc.subjectRepo.CreateSubject(ctx, info); err != nil {
		logger.Error("Fail to create subjects")
		return err
	}

	if err := sc.subjectFieldRepo.CreateSubjectField(
		ctx, subjectId, req.CreateSubject.Year, req.CreateSubject.Month, req.CreateSubject.Date,
		req.CreateSubject.WeekDay,
	); err != nil {
		logger.Error("Failed to create subject field")
		return err
	}
	return nil
}

func (sc SubjectCtrl) saveFile(file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	// ignore
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)
	dst := sc.cfg.StaticDirectory + "/" + fileName
	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	// ignore
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = io.Copy(out, src)
	return err
}
