package user

import (
	"backend/ent"
	"backend/ent/members"
	"backend/internal/logger"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m MysqlRepo) CreateNewUser(ctx context.Context, members *ent.Members) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(members.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to encrypt password")
		return err
	}
	_, err = m.client.Members.Create().
		SetUsername(members.Username).SetNickname(members.Nickname).
		SetEmail(members.Email).SetPassword(string(encryptedPassword)).
		SetRegisterTime(members.RegisterTime).Save(ctx)
	if err != nil {
		logger.Error("Failed to create member")
		return err
	}
	return nil
}

func (m MysqlRepo) GetByUid(ctx context.Context, uid uint32) (member *ent.Members, err error) {
	member, err = m.client.Members.Query().Where(members.IDEQ(uid)).Only(ctx)
	return
}

func (m MysqlRepo) GetByEmail(ctx context.Context, email string) (member *ent.Members, err error) {
	member, err = m.client.Members.Query().Where(members.EmailEQ(email)).Only(ctx)
	return
}

func (m MysqlRepo) DeleteByUid(ctx context.Context, uid uint32) error {
	err := m.client.Members.DeleteOneID(uid).Exec(ctx)
	if err != nil {
		logger.Error("Failed to delete user")
		return err
	}
	return nil
}
