package subject

import (
	"backend/ent"
	"backend/ent/members"
	"backend/ent/subject"
	"backend/internal/logger"
	"backend/internal/model"
	"backend/web/request/collection"
	"context"
	"errors"
	"time"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}
func (m MysqlRepo) GetSubject(ctx context.Context) (ent.Subjects, error) {
	subjects, err := m.client.Subject.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (m MysqlRepo) GetSubjectByID(ctx context.Context, id int64) (ent.Subject, error) {
	subject, err := m.client.Subject.Query().Where(subject.ID(int(id))).First(ctx)
	if err != nil {
		return *subject, err
	}
	return *subject, nil
}

func (m MysqlRepo) SearchSubjection(ctx context.Context, name string) (*ent.Subject, error) {
	Subject, err := m.client.Subject.Query().Where(subject.Name(name)).First(ctx)
	if err != nil {
		return Subject, err
	}
	return Subject, nil
}
func (m MysqlRepo) UpdateCollection(ctx context.Context, MemberId uint32, SubjectId int, req collection.UpdateCollectionReq) error {
	//mem, err := m.client.Members.Query().Where(members.ID(MemberId)).First(ctx)
	//mem.QueryCollections().s
	//col, _ := m.client.Members.QueryCollections(mem).All(ctx)
	//m.client.Collection.Query
	//m.client..Query().Where(c2.HasSubjectWith())

	//if err != nil {
	//	return err
	//}
	//if req.Comment != "" {
	//	col.Update().SetType(req.CollectionType).SetIfComment(true).SetComment(req.Comment).
	//		SetTime(time.Now().String()).SetScore(req.Score).Save(ctx)
	//}

	//col.Update().AddType(int8(req.CollectionType)).SetScore(req.Score).SetTime(time.Now().String()).Save(ctx)
	return nil
}
func (m MysqlRepo) AddOrUpdateCollection(ctx context.Context, MemberId uint32, SubjectId int, AddType model.CollectionType) error {
	member, _ := m.client.Members.Query().Where(members.ID(MemberId)).First(ctx)
	sub, _ := m.client.Subject.Query().Where(subject.ID(SubjectId)).First(ctx)
	_, err := m.client.Collection.Create().SetMember(member).SetSubject(sub).SetType(AddType).SetTime(time.Now().String()).Save(ctx)

	if err != nil {
		return err
	}
	switch AddType {
	case 1:
		sub.Update().AddCollect(1).AddWish(1).Save(ctx)
	case 2:
		sub.Update().AddCollect(1).AddOnHold(1).Save(ctx)
	case 3:
		sub.Update().AddCollect(1).AddDoing(1).Save(ctx)
	case 4:
		sub.Update().AddCollect(1).AddDrop(1).Save(ctx)
	case 5:
		sub.Update().AddCollect(1).AddWatched(1).Save(ctx)
	default:
		return errors.New("unknown collection type")
	}
	sub.Update().AddCollect(1).AddWish(1).Save(ctx)
	return nil
}
func (m MysqlRepo) CreateSubject(ctx context.Context, subject1 *ent.Subject) error {
	if u, _ := m.client.Subject.Query().Where(subject.Name(subject1.Name)).First(ctx); u != nil {
		logger.Error("created same subject")
		err := errors.New("already created same subject")
		return err
	}
	_, err := m.client.Subject.Create().
		SetName(subject1.Name).
		SetNameCn(subject1.NameCn).SetSubjectType(subject1.SubjectType).
		SetDate(subject1.Date).SetSummary(subject1.Summary).
		SetImage(subject1.Image).Save(ctx)
	if err != nil {
		logger.Error("error to create this subject")
		return err
	}
	return nil
}
