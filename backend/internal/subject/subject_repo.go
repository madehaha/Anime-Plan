package subject

import (
	"context"
	"errors"
	"time"

	"backend/ent"
	"backend/ent/subject"
	"backend/ent/subjectfield"
	"backend/internal/collection"
	"backend/internal/logger"
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

func (m MysqlRepo) GetSubjectByID(ctx context.Context, subjectId uint32) (*ent.Subject, *ent.SubjectField, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.ID(subjectId)).First(ctx)
	field, err := m.client.SubjectField.Query().Where(subjectfield.HasSubjectWith(subject.ID(subjectId))).First(ctx)
	if err != nil {
		return subjectEntity, nil, err
	}
	return subjectEntity, field, nil
}

func (m MysqlRepo) Rankings(ctx context.Context) ([]Middle, error) {
	subjects, err := m.client.Subject.Query().Order(subject.BySubjectFieldField(subjectfield.FieldRank)).All(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("no ranks")
	}
	length := len(subjects)
	res := make([]Middle, length)
	for key, value := range subjects {
		res[key].Subject = new(ent.Subject)
		*res[key].Subject = *value
		res[key].Field = new(ent.SubjectField)
		mid, _ := subjects[key].QuerySubjectField().First(ctx)
		*res[key].Field = *mid
	}

	return res, nil
}

func (m MysqlRepo) BoardCast(ctx context.Context, Day uint8) (ent.Subjects, error) {
	year := uint32(time.Now().Year())
	month := uint8(time.Now().Month())
	println(month)
	subjectEntity, err := m.client.Subject.Query().Where(
		subject.And(
			subject.HasSubjectFieldWith(subjectfield.Year(year)),
			subject.HasSubjectFieldWith(subjectfield.Weekday(Day)),
			subject.HasSubjectFieldWith(subjectfield.MonthGTE(month+3)),
			subject.HasSubjectFieldWith(subjectfield.MonthLTE(month-3)),
		),
	).All(ctx)
	if err != nil {
		return subjectEntity, err
	}
	return subjectEntity, nil
}

func (m MysqlRepo) GetSubjectByName(ctx context.Context, name string) (*ent.Subject, *ent.SubjectField, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.Name(name)).First(ctx)
	if err != nil {
		return subjectEntity, nil, err
	}
	field, err := m.client.SubjectField.Query().Where(subjectfield.HasSubjectWith(subject.Name(name))).First(ctx)
	if err != nil {
		return subjectEntity, nil, err
	}
	return subjectEntity, field, nil
}
func (m MysqlRepo) GetSubjectByNameCN(ctx context.Context, name string) (*ent.Subject, *ent.SubjectField, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.NameCn(name)).First(ctx)
	if err != nil {
		return subjectEntity, nil, err
	}
	field, err := m.client.SubjectField.Query().Where(subjectfield.HasSubjectWith(subject.NameCn(name))).First(ctx)
	if err != nil {
		return subjectEntity, nil, err
	}
	return subjectEntity, field, nil
}

func (m MysqlRepo) CreateSubject(ctx context.Context, info InitialInfo) (uint32, error) {
	if u, _ := m.client.Subject.Query().Where(subject.Name(info.Name)).First(ctx); u != nil {
		logger.Error("created same subject")
		err := errors.New("already created same subject")
		return 0, err
	}
	subjectEntity, err := m.client.Subject.Create().SetImage(info.Image).SetSummary(info.Summary).SetName(info.Name).
		SetNameCn(info.NameCN).SetEpisodes(info.Episodes).Save(ctx)
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
	flag, err := m.client.Subject.Query().Where(subject.ID(subjectId)).Exist(ctx)
	if !flag {
		return err
	}

	err = m.addSubjectType(ctx, subjectId, newType)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	previousNumber, err := m.getTypeNumber(ctx, subjectId, previousType)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = m.minusSubjectType(ctx, subjectId, previousNumber, previousType)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (m MysqlRepo) getTypeNumber(
	ctx context.Context, subjectId uint32, collectionType collection.CollectionType,
) (uint32, error) {
	subjectEntity, err := m.client.Subject.Query().Where(subject.ID(subjectId)).Only(ctx)
	if err != nil {
		logger.Error("Failed to get subject by subject id")
		return 0, err
	}
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
