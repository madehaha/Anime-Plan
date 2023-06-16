package collection

import (
	"context"
	"fmt"
	"time"

	"backend/ent"
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/subject"
	"backend/internal/logger"
	"backend/internal/model"
	collectionReq "backend/web/request/collection"
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
	ctx context.Context, uid uint32, subjectId uint32, req collectionReq.AddOrUpdateCollectionReq,
) error {
	var score uint8
	var epStatus uint8
	date := time.Now().Format("2006-01-02")
	hasComment := req.Comment != ""
	hasEpStatusOrScore := req.Type == model.Wish
	if hasEpStatusOrScore {
		score = req.Score
		epStatus = req.EpStatus
	} else {
		score = 0
		epStatus = 0
	}
	// not exist
	_, err := m.Client.Collection.Create().SetType(req.Type).SetHasComment(hasComment).SetComment(req.Comment).
		SetScore(score).SetAddTime(date).SetEpStatus(epStatus).SetMemberID(uid).SetSubjectID(subjectId).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) UpdateCollection(
	ctx context.Context, collectionEntity *ent.Collection,
) error {

	err := m.Client.Collection.UpdateOne(collectionEntity).SetType(collectionEntity.Type).SetHasComment(
		collectionEntity.HasComment,
	).SetComment(collectionEntity.Comment).SetScore(collectionEntity.Score).
		SetAddTime(collectionEntity.AddTime).SetEpStatus(collectionEntity.EpStatus).Exec(ctx)

	return err
}

func (m MysqlRepo) DeleteCollection(ctx context.Context, uid uint32, subjectId uint32) error {
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
