package ctrl

import (
	"context"
	"errors"
	"time"

	"backend/ent"

	"backend/internal/collection"
	"backend/internal/logger"
	"backend/internal/subject"
	"backend/internal/subject_field"
	collectionReq "backend/web/request/collection"
)

type CollectionCtrl struct {
	collectionRepo   collection.MysqlRepo
	subjectRepo      subject.MysqlRepo
	subjectFieldRepo subjectField.MysqlRepo
}

func NewCollectionCtrl(
	collectionRepo collection.MysqlRepo, subjectRepo subject.MysqlRepo,
	subjectFieldRepo subjectField.MysqlRepo,
) CollectionCtrl {
	return CollectionCtrl{
		collectionRepo:   collectionRepo,
		subjectRepo:      subjectRepo,
		subjectFieldRepo: subjectFieldRepo,
	}
}

func (cc CollectionCtrl) AddCollection(
	uid uint32, subjectId uint32, req collectionReq.AddOrUpdateCollectionReq,
) error {
	ctx := context.Background()

	if collectionEntity, _ := cc.collectionRepo.GetCollectionByUidAndSubjectId(
		ctx, uid, subjectId,
	); collectionEntity != nil {
		// the collection already exist
		logger.Error("Trying to add duplicate collection")
		return errors.New("collection is already existed")
	}

	collectionInfo := collection.Info{
		Type:       req.Type,
		HasComment: *req.HasComment,
		Comment:    req.Comment,
		Score:      req.Score,
		AddTime:    time.Now().Format("2006-01-02"),
		EpStatus:   req.EpStatus,
	}

	err := cc.collectionRepo.AddCollection(ctx, uid, subjectId, collectionInfo)
	if err != nil {
		logger.Error("Failed to add collection")
		return err
	}
	err = cc.subjectRepo.AddSubjectType(ctx, subjectId, req.Type)
	if err != nil {
		logger.Error("Failed to add subject type")
		return err
	}

	err = cc.subjectFieldRepo.UpdateSubjectField(ctx, subjectId, 0, req.Score)
	if err != nil {
		logger.Error("Failed to add subject field")
	}

	return nil
}

func (cc CollectionCtrl) UpdateCollection(
	uid uint32, subjectId uint32, req collectionReq.AddOrUpdateCollectionReq,
) error {
	// TODO rewrite
	// 1. check if exist, if not error 			DONE
	// 2. update collection 					DONE
	// 3. update subject
	// 4. update subject field
	// 5. update subject rank
	ctx := context.Background()
	collectionEntity, _ := cc.collectionRepo.GetCollectionByUidAndSubjectId(ctx, uid, subjectId)
	if collectionEntity == nil {
		logger.Error("the collection not exist.")
		return errors.New("the collection not exist")
	}

	collectionInfo := collection.Info{
		Type:       req.Type,
		HasComment: *req.HasComment,
		Comment:    req.Comment,
		Score:      req.Score,
		AddTime:    time.Now().Format("2006-01-02"),
		EpStatus:   req.EpStatus,
	}

	// update collection
	err := cc.collectionRepo.UpdateCollection(ctx, collectionEntity.ID, collectionInfo)
	if err != nil {
		logger.Error("Failed to update collection")
		return err
	}

	// update subject
	err = cc.subjectRepo.UpdateSubjectType(ctx, subjectId, collectionEntity.Type, req.Type)
	if err != nil {
		logger.Error("Failed to update subject")
		return err
	}

	// update subject field
	err = cc.subjectFieldRepo.UpdateSubjectField(ctx, subjectId, collectionEntity.Score, req.Score)
	if err != nil {
		logger.Error("Failed to update subject field")
		return err
	}
	return nil

	// update subject
	// ctx := context.Background()
	// collectionEntity, err := cc.collectionRepo.GetCollectionByUidAndSubjectId(ctx, uid, subjectId)
	// if err != nil {
	// 	return err
	// }
	// logger.Info(collectionEntity.String())
	// // update
	// hasComment := req.Comment != ""
	// collectionEntity.HasComment = hasComment
	// hasEpStatusOrScore := req.Type != collection.Wish
	// if hasEpStatusOrScore {
	// 	collectionEntity.EpStatus = req.EpStatus
	// 	collectionEntity.Score = req.Score
	// } else {
	// 	collectionEntity.EpStatus = 0
	// 	collectionEntity.Score = 0
	// }
	// collectionEntity.Comment = req.Comment
	// collectionEntity.Type = req.Type
	// collectionEntity.AddTime = time.Now().Format("2006-01-02")
	// logger.Info(collectionEntity.String())
	// err = cc.collectionRepo.UpdateCollection(ctx, collectionEntity)
	// if err != nil {
	// 	logger.Error("Failed to update collection")
	// 	return err
	// }
	// return err
}

func (cc CollectionCtrl) DeleteCollection(uid uint32, subjectId uint32) error {
	return cc.collectionRepo.DeleteCollection(context.Background(), uid, subjectId)
}

func (cc CollectionCtrl) GetCollectionById(searchType string, id uint32) (ent.Collections, error) {
	collections, err := cc.collectionRepo.GetCollectionByID(context.Background(), searchType, id)
	if err != nil {
		return nil, err
	}
	return collections, err
}

func (cc CollectionCtrl) GetSelfCollection(subject_id uint32, member_id uint32) (*ent.Collection, error) {
	Collection, err := cc.collectionRepo.GetCollectionByUidAndSubjectId(context.Background(), member_id, subject_id)
	if err != nil {
		return Collection, err
	}
	return Collection, nil
}
func (cc CollectionCtrl) GetSelfCollections(member_id uint32, Type uint8) (ent.Collections, error) {
	Collection, err := cc.collectionRepo.GetCollectionByUidAndType(context.Background(), member_id, Type)
	if err != nil {
		return Collection, err
	}
	return Collection, nil
}
