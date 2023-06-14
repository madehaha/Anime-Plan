package ctrl

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/logger"
	"backend/internal/user"
	userReq "backend/web/request/user"
	"backend/web/util"
)

type UserCtrl struct {
	Repo user.MysqlRepo
	Jwt  util.JwtUtil
}

func NewUserCtrl(repo user.MysqlRepo, jwtUtil util.JwtUtil) UserCtrl {
	return UserCtrl{
		Repo: repo,
		Jwt:  jwtUtil,
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
