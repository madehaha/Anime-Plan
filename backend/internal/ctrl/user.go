package ctrl

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/internal/user"
	"backend/web/util"
	"context"
	"golang.org/x/crypto/bcrypt"
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

func (uc UserCtrl) Register(members *ent.Members) error {
	if err := uc.Repo.CreateNewUser(context.Background(), members); err != nil {
		logger.Error("Failed to create new user")
		return err
	}
	return nil
}

func (uc UserCtrl) Login(email string, password string) (token string, err error) {
	member, err := uc.Repo.GetByEmail(context.Background(), email)
	if err != nil {
		logger.Error("Failed to get member by email")
		return
	}
	encryptedPassword := member.Password
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		// wrong password
		logger.Info("Wrong password")
		return
	}
	token = uc.Jwt.GenerateNewToken(member.ID, member.Gid)
	return
}
