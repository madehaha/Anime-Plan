package collection

import (
	"context"
	"fmt"

	"backend/ent"
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/subject"
	"backend/internal/logger"
)

type MysqlRepo struct {
	Client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{Client: client}
}

func (m MysqlRepo) GetCollectionByUidAndSubjectId(
	ctx context.Context, uid uint32,
	subjectId uint32,
) (collectionEntity *ent.Collection, err error) {
	collectionEntity, err = m.Client.Collection.Query().Where(
		collection.And(
			collection.HasSubjectWith(subject.ID(subjectId)),
			collection.HasMemberWith(members.ID(uid)),
		),
	).Only(ctx)
	return
}

func (m MysqlRepo) AddCollection(
	ctx context.Context, uid uint32, subjectId uint32, collectionInfo Info,
) error {
	if _, err := m.Client.Collection.Create().SetType(collectionInfo.Type).SetHasComment(collectionInfo.HasComment).
		SetComment(collectionInfo.Comment).SetScore(collectionInfo.Score).SetAddTime(collectionInfo.AddTime).SetEpStatus(collectionInfo.EpStatus).SetMemberID(uid).
		SetSubjectID(subjectId).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) UpdateCollection(
	ctx context.Context, collectionId uint32, info Info,
) error {
	err := m.Client.Collection.UpdateOneID(collectionId).SetType(info.Type).SetHasComment(info.HasComment).
		SetComment(info.Comment).SetScore(info.Score).SetAddTime(info.AddTime).SetEpStatus(info.EpStatus).Exec(ctx)
	return err
}

func (m MysqlRepo) GetCollectionByID(ctx context.Context, searchType string, id uint32) (ent.Collections, error) {
	switch searchType {
	case "subject":
		{
			collections, err := m.Client.Collection.Query().Where(collection.HasSubjectWith(subject.ID(id))).All(ctx)
			return collections, err
		}
	case "member":
		{
			collections, err := m.Client.Collection.Query().Where(collection.HasMemberWith(members.ID(id))).All(ctx)
			return collections, err
		}
	default:
		return nil, nil
	}
}

func (m MysqlRepo) DeleteCollection(ctx context.Context, uid uint32, subjectId uint32) error {
	// TODO
	number, err := m.Client.Collection.Delete().Where(
		collection.And(
			collection.HasMemberWith(members.ID(uid)),
			collection.HasSubjectWith(subject.ID(subjectId)),
		),
	).Exec(ctx)
	s := fmt.Sprintf("Delete collection %d", number)
	logger.Info(s)
	return err
}
