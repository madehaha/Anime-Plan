package user

import (
	"backend/ent"
	"backend/internal/logger"
	"context"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m MysqlRepo) CreateNewUser(ctx context.Context, members *ent.Members) error {
	_, err := m.client.Members.Create().
		SetUsername(members.Username).SetNickname(members.Nickname).
		SetEmail(members.Email).SetPassword(members.Password).
		SetRegisterTime(members.RegisterTime).Save(ctx)
	if err != nil {
		logger.Error("Failed to create member")
		return err
	}
	return nil
}
