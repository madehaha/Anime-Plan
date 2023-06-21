package subjectField

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"backend/ent"
	"backend/ent/subject"
	"backend/ent/subjectfield"
	"backend/internal/logger"
	subjectModel "backend/internal/subject"
)

type MysqlRepo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) MysqlRepo {
	return MysqlRepo{client: client}
}

func (m MysqlRepo) CreateSubjectField(
	ctx context.Context, subjectId uint32, year uint32, month uint8,
	date uint8, weekday uint8,
) error {
	err := m.client.SubjectField.Create().SetSubjectID(subjectId).SetYear(year).
		SetMonth(month).SetDate(date).SetWeekday(weekday).Exec(ctx)
	return err
}

func (m MysqlRepo) GetSubjectFieldBySubjectId(ctx context.Context, subjectId uint32) (*ent.SubjectField, error) {
	return m.client.SubjectField.Query().Where(subjectfield.HasSubjectWith(subject.ID(subjectId))).Only(ctx)
}

func (m MysqlRepo) UpdateSubjectField(ctx context.Context, subjectId uint32, previousScore uint8, score uint8) error {
	if err := m.UpdateSubjectFieldRate(ctx, subjectId, previousScore, score); err != nil {
		logger.Error("Failed to update subject field")
		return err
	}
	subjectEntity, err := m.GetSubjectFieldBySubjectId(ctx, subjectId)
	if err != nil {
		logger.Error("Failed to get subject field by subject id")
		return err
	}
	// update score
	if err := m.updateSubjectFieldScore(ctx, subjectEntity); err != nil {
		logger.Error("Failed to get subject field by subject id")
		return err
	}
	// update rank
	err = m.updateSubjectFieldsRank(ctx)
	if err != nil {
		logger.Error("Failed to update rank")
		return err
	}
	return nil
}

func (m MysqlRepo) UpdateSubjectFieldRate(
	ctx context.Context, subjectId uint32, previousScore uint8, score uint8,
) error {
	var err error
	var previousNumber uint32
	if score == previousScore {
		return nil
	}
	s := fmt.Sprintf("previous: %d score: %d", previousScore, score)
	logger.Info(s)

	if previousScore == 0 {
		err = m.addRateX(ctx, subjectId, score)
	} else {
		if previousNumber, err = m.getRateX(ctx, subjectId, previousScore); err != nil {
			logger.Error("Failed to get the previous number of score")
			return err
		}
		if score == 0 {
			err = m.minusRateX(ctx, subjectId, previousScore, previousNumber)
		} else {
			err = m.addRateX(ctx, subjectId, score)
			err = m.minusRateX(ctx, subjectId, previousScore, previousNumber)
		}
	}

	return err
}

func (m MysqlRepo) getRateX(ctx context.Context, subjectId uint32, score uint8) (uint32, error) {
	entity, err := m.client.SubjectField.Query().Where(subjectfield.HasSubjectWith(subject.ID(subjectId))).Only(ctx)
	if err != nil {
		logger.Error("Failed to get rate")
		return 0, err
	}
	switch score {
	case 1:
		return entity.Rate1, nil
	case 2:
		return entity.Rate2, nil
	case 3:
		return entity.Rate3, nil
	case 4:
		return entity.Rate4, nil
	case 5:
		return entity.Rate5, nil
	case 6:
		return entity.Rate6, nil
	case 7:
		return entity.Rate7, nil
	case 8:
		return entity.Rate8, nil
	case 9:
		return entity.Rate9, nil
	case 10:
		return entity.Rate10, nil
	default:
		return 0, errors.New("wrong score")
	}
}

func (m MysqlRepo) addRateX(ctx context.Context, subjectId uint32, score uint8) error {
	updateOne := m.client.SubjectField.Update().Where(subjectfield.HasSubjectWith(subject.ID(subjectId)))
	switch score {
	case 1:
		return updateOne.AddRate1(1).Exec(ctx)
	case 2:
		return updateOne.AddRate2(1).Exec(ctx)
	case 3:
		return updateOne.AddRate3(1).Exec(ctx)
	case 4:
		return updateOne.AddRate4(1).Exec(ctx)
	case 5:
		return updateOne.AddRate5(1).Exec(ctx)
	case 6:
		return updateOne.AddRate6(1).Exec(ctx)
	case 7:
		return updateOne.AddRate7(1).Exec(ctx)
	case 8:
		return updateOne.AddRate8(1).Exec(ctx)
	case 9:
		return updateOne.AddRate9(1).Exec(ctx)
	case 10:
		return updateOne.AddRate10(1).Exec(ctx)
	default:
		return errors.New("wrong score")
	}
}

func (m MysqlRepo) minusRateX(ctx context.Context, subjectId uint32, previousScore uint8, previousNumber uint32) error {
	updateOne := m.client.SubjectField.Update().Where(subjectfield.HasSubjectWith(subject.ID(subjectId)))
	switch previousScore {
	case 1:
		return updateOne.SetRate1(previousNumber - 1).Exec(ctx)
	case 2:
		return updateOne.SetRate2(previousNumber - 1).Exec(ctx)
	case 3:
		return updateOne.SetRate3(previousNumber - 1).Exec(ctx)
	case 4:
		return updateOne.SetRate4(previousNumber - 1).Exec(ctx)
	case 5:
		return updateOne.SetRate5(previousNumber - 1).Exec(ctx)
	case 6:
		return updateOne.SetRate6(previousNumber - 1).Exec(ctx)
	case 7:
		return updateOne.SetRate7(previousNumber - 1).Exec(ctx)
	case 8:
		return updateOne.SetRate8(previousNumber - 1).Exec(ctx)
	case 9:
		return updateOne.SetRate9(previousNumber - 1).Exec(ctx)
	case 10:
		return updateOne.SetRate10(previousNumber - 1).Exec(ctx)
	default:
		return errors.New("wrong previousScore")
	}
}

func (m MysqlRepo) updateSubjectFieldScore(ctx context.Context, subjectField *ent.SubjectField) error {
	var score float64
	var number uint32 = subjectField.Rate1 + subjectField.Rate2 + subjectField.Rate3 + subjectField.Rate4 + subjectField.Rate5 +
		subjectField.Rate6 + subjectField.Rate7 + subjectField.Rate8 + subjectField.Rate9 + subjectField.Rate10
	if number == 0 {
		score = 0
	} else {
		score = float64(
			subjectField.Rate1*1+subjectField.Rate2*2+subjectField.Rate3*3+subjectField.Rate4*4+
				subjectField.Rate5*5+subjectField.Rate6*6+subjectField.Rate7*7+subjectField.Rate8*8+
				subjectField.Rate9*9+subjectField.Rate10*10,
		) / float64(number)
	}
	return m.client.SubjectField.UpdateOne(subjectField).SetAverageScore(score).Exec(ctx)
}

func (m MysqlRepo) updateSubjectFieldsRank(ctx context.Context) error {
	subjectFieldEntities, err := m.client.SubjectField.Query().Order(subjectfield.ByAverageScore(sql.OrderDesc())).All(ctx)
	if err != nil {
		return err
	}
	for i, entity := range subjectFieldEntities {
		err := m.updateSubjectFieldRank(ctx, uint32(i+1), entity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m MysqlRepo) updateSubjectFieldRank(ctx context.Context, rank uint32, field *ent.SubjectField) error {
	return m.client.SubjectField.UpdateOne(field).SetRank(rank).Exec(ctx)
}

func (m MysqlRepo) InsertSubjectField(ctx context.Context, subjectId uint32, subjectModel subjectModel.Subject) error {
	return m.client.SubjectField.Create().SetRate1(subjectModel.Rate1).SetRate2(subjectModel.Rate2).
		SetRate3(subjectModel.Rate3).SetRate4(subjectModel.Rate4).
		SetRate5(subjectModel.Rate5).SetRate6(subjectModel.Rate6).
		SetRate7(subjectModel.Rate7).SetRate8(subjectModel.Rate8).
		SetRate9(subjectModel.Rate9).SetRate10(subjectModel.Rate10).
		SetAverageScore(subjectModel.AverageScore).SetYear(subjectModel.Year).
		SetMonth(subjectModel.Month).SetDate(subjectModel.Date).
		SetWeekday(subjectModel.Weekday).SetSubjectID(subjectId).Exec(ctx)
}
