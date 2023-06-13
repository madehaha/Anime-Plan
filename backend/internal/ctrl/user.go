package ctrl

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/internal/user"
	"context"
)

type UserCtrl struct {
	Repo user.MysqlRepo
}

func NewUserCtrl(repo user.MysqlRepo) UserCtrl {
	return UserCtrl{Repo: repo}
}

func (uc UserCtrl) Register(members *ent.Members) error {
	if err := uc.Repo.CreateNewUser(context.Background(), members); err != nil {
		logger.Error("Failed to create new user")
		return err
	}
	return nil
}
