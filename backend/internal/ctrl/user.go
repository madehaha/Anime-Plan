package ctrl

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/user"
	userReq "backend/web/request/user"
	"backend/web/util"
)

type UserCtrl struct {
	Repo user.MysqlRepo
	Jwt  util.JwtUtil
	cfg  config.AppConfig
}

func NewUserCtrl(
	repo user.MysqlRepo, jwtUtil util.JwtUtil,
	cfg config.AppConfig,
) UserCtrl {
	return UserCtrl{
		Repo: repo,
		Jwt:  jwtUtil,
		cfg:  cfg,
	}
}

func (uc UserCtrl) Register(r userReq.UserRegisterReq) error {
	if err := uc.Repo.CreateNewUser(context.Background(), r); err != nil {
		logger.Error("Failed to create new user")
		return err
	}
	return nil
}

func (uc UserCtrl) Login(email string, password string) (
	token string, err error,
) {
	member, err := uc.Repo.GetByEmail(context.Background(), email)
	if err != nil {
		logger.Error("Failed to get member by email")
		return
	}
	encryptedPassword := member.Password
	err = bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword), []byte(password),
	)
	if err != nil {
		// wrong password
		logger.Info("Wrong password")
		return
	}
	token = uc.Jwt.GenerateNewToken(member.ID, member.Gid)
	return
}

func (uc UserCtrl) Cancel(uid uint32) error {
	err := uc.Repo.DeleteByUid(context.Background(), uid)
	if err != nil {
		logger.Error("Failed to delete by uid")
		return err
	}
	return nil
}

func (uc UserCtrl) GetAvtar(uid uint32) (string, error) {
	member, err := uc.Repo.GetByUid(context.Background(), uid)
	if err != nil {
		logger.Error("Failed to get user by uid")
		return "", err
	}
	return member.Avatar, nil
}

func (uc UserCtrl) ModifyInfo(uid uint32, r userReq.UserModifyInfoResp) error {
	// write file
	fileHeader := r.FileData
	filename := strconv.FormatUint(uint64(uid), 10) + "_" + time.Now().Format(
		"200601021504",
	) + "_" + filepath.Ext(
		fileHeader.Filename,
	)
	err := uc.saveFile(fileHeader, filename)
	if err != nil {
		logger.Error("Failed to save file")
		return err
	}
	imageUrl := uc.cfg.WebDomain + "/" + uc.cfg.StaticDirectory + "/" + filename
	logger.Info(imageUrl)
	err = uc.Repo.UpdateByUid(context.Background(), uid, r.Info, imageUrl)
	if err != nil {
		logger.Error("Failed to update")
		return err
	}
	return nil
}

func (uc UserCtrl) saveFile(file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	// ignore
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)
	dst := uc.cfg.StaticDirectory + "/" + fileName
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
