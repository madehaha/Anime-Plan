package ctrl

import (
	"context"
	"time"

	"backend/internal/collection"
	"backend/internal/logger"
	"backend/internal/model"
	collectionReq "backend/web/request/collection"
)

type CollectionCtrl struct {
	Repo collection.MysqlRepo
}

func NewCollectionCtrl(repo collection.MysqlRepo) CollectionCtrl {
	return CollectionCtrl{
		Repo: repo,
	}
}

func (cc CollectionCtrl) AddCollection(
	uid uint32, subjectId uint32, req collectionReq.AddOrUpdateCollectionReq,
) error {
	ctx := context.Background()
	err := cc.Repo.AddCollection(ctx, uid, subjectId, req)
	if err != nil {
		logger.Error("Failed to add collection")
		return err
	}
	return nil
}

func (cc CollectionCtrl) UpdateCollection(
	uid uint32, subjectId uint32, req collectionReq.AddOrUpdateCollectionReq,
) error {
	ctx := context.Background()
	collectionEntity, err := cc.Repo.GetCollectionByUidAndSubjectId(ctx, uid, subjectId)
	if err != nil {
		return err
	}
	// update
	hasComment := req.Comment != ""
	collectionEntity.HasComment = hasComment
	hasEpStatusOrScore := req.Type != model.Wish
	if hasEpStatusOrScore {
		collectionEntity.EpStatus = req.EpStatus
		collectionEntity.Score = req.Score
	} else {
		collectionEntity.EpStatus = 0
		collectionEntity.Score = 0
	}
	collectionEntity.Type = req.Type
	collectionEntity.AddTime = time.Now().Format("2006-01-02")
	err = cc.Repo.UpdateCollection(ctx, collectionEntity)
	if err != nil {
		logger.Error("Failed to update collection")
		return err
	}
	return err
}
