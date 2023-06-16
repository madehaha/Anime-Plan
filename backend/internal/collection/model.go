package collection

import (
	"context"
	"time"

	"backend/ent"
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/subject"
	"backend/internal/model"
	collectionReq "backend/web/request/collection"
)

type MysqlRepo struct {
	Client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{Client: client}
}

// func (m MysqlRepo) UpdateCollection(
//
//	ctx context.Context, MemberId uint32, SubjectId int, req collectionReq.UpdateCollectionReq,
//
//	) error {
//		// mem, err := m.client.Members.Query().Where(members.ID(MemberId)).First(ctx)
//		// col, _ := m.client.Members.QueryCollections(mem).All(ctx)
//		// m.client.Collection.Query
//		// col, _ := mem.QueryCollections().QuerySubject()
//		if err != nil {
//			return err
//		}
//		ent.Collection{}
//		m.client.Subject.Query().Where(subject.ID(req.))
//		m.client.Members.Query().Where(members.Has)
//		if req.Comment != "" {
//			col.Update().SetType(req.CollectionType).SetIfComment(true).SetComment(req.Comment).
//				SetTime(time.Now().String()).SetScore(req.Score).Save(ctx)
//		}
//
//		col.Update().AddType(int8(req.CollectionType)).SetScore(req.Score).SetTime(time.Now().String()).Save(ctx)
//		return nil
//	}
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
	_, err := m.Client.Collection.UpdateOne(collectionEntity).Save(ctx)
	return err
}
