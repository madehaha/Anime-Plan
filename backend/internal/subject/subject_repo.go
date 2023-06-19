package subject

import (
	"context"
	"errors"

	"backend/ent"
	"backend/ent/subject"
	"backend/internal/collection"
	"backend/internal/logger"
	subjectReq "backend/web/request/subject"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m MysqlRepo) GetSubject(ctx context.Context) (ent.Subjects, error) {
	subjectEntities, err := m.client.Subject.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return subjectEntities, nil
}

func (m MysqlRepo) GetSubjectByID(ctx context.Context, subjectId uint32) (*ent.Subject, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.ID(subjectId)).First(ctx)
	if err != nil {
		return subjectEntity, err
	}
	return subjectEntity, nil
}

func (m MysqlRepo) GetSubjectByName(ctx context.Context, name string) (*ent.Subject, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.Name(name)).First(ctx)
	if err != nil {
		return subjectEntity, err
	}
	return subjectEntity, nil
}

func (m MysqlRepo) CreateSubject(ctx context.Context, req subjectReq.CreateSubjectReq) (uint32, error) {
	if u, _ := m.client.Subject.Query().Where(subject.Name(req.Name)).First(ctx); u != nil {
		logger.Error("created same subject")
		err := errors.New("already created same subject")
		return 0, err
	}
	subjectEntity, err := m.client.Subject.Create().SetImage(req.Image).SetSummary(req.Summary).SetName(req.Name).
		SetNameCn(req.NameCN).SetEpisodes(req.Episodes).Save(ctx)
	if err != nil {
		logger.Error("error to create the subject")
		return 0, err
	}

	return subjectEntity.ID, nil
}

func (m MysqlRepo) AddSubjectType(
	ctx context.Context, subjectId uint32, collectionType collection.CollectionType,
) error {
	return m.addSubjectType(ctx, subjectId, collectionType)
}

func (m MysqlRepo) UpdateSubjectType(
	ctx context.Context, subjectId uint32, previousType collection.CollectionType, newType collection.CollectionType,
) error {
	subjectEntity, err := m.client.Subject.Query().Where(subject.ID(subjectId)).Only(ctx)
	if err != nil {
		return err
	}
	previousNumber, err := getTypeNumber(subjectEntity, previousType)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = m.addSubjectType(ctx, subjectId, newType)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = m.minusSubjectType(ctx, subjectId, previousNumber, previousType)
	return err
}

func getTypeNumber(subjectEntity *ent.Subject, collectionType collection.CollectionType) (uint32, error) {
	switch collectionType {
	case collection.Wish:
		return subjectEntity.Wish, nil
	case collection.Watched:
		return subjectEntity.Watched, nil
	case collection.Doing:
		return subjectEntity.Doing, nil
	case collection.OnHold:
		return subjectEntity.OnHold, nil
	case collection.Dropped:
		return subjectEntity.Dropped, nil
	default:
		return 0, errors.New("wrong collection type")
	}
}

func (m MysqlRepo) addSubjectType(
	ctx context.Context, subjectId uint32,
	collectionType collection.CollectionType,
) error {
	subjectUpdateOne := m.client.Subject.UpdateOneID(subjectId)
	switch collectionType {
	case collection.Wish:
		return subjectUpdateOne.AddWish(1).Exec(ctx)
	case collection.Watched:
		return subjectUpdateOne.AddWatched(1).Exec(ctx)
	case collection.Doing:
		return subjectUpdateOne.AddDoing(1).Exec(ctx)
	case collection.OnHold:
		return subjectUpdateOne.AddOnHold(1).Exec(ctx)
	case collection.Dropped:
		return subjectUpdateOne.AddDropped(1).Exec(ctx)
	default:
		return errors.New("fucked up, you passed the wrong collection type")
	}
}

func (m MysqlRepo) minusSubjectType(
	ctx context.Context, subjectId uint32, previousNumber uint32,
	collectionType collection.CollectionType,
) error {
	subjectUpdateOne := m.client.Subject.UpdateOneID(subjectId)
	switch collectionType {
	case collection.Wish:
		return subjectUpdateOne.SetWish(previousNumber - 1).Exec(ctx)
	case collection.Watched:
		return subjectUpdateOne.SetWatched(previousNumber - 1).Exec(ctx)
	case collection.Doing:
		return subjectUpdateOne.SetDoing(previousNumber - 1).Exec(ctx)
	case collection.OnHold:
		return subjectUpdateOne.SetOnHold(previousNumber - 1).Exec(ctx)
	case collection.Dropped:
		return subjectUpdateOne.SetDropped(previousNumber - 1).Exec(ctx)
	default:
		return errors.New("fucked up, you passed the wrong collection type")
	}
}
