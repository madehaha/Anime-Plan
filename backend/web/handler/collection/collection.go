package collection

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/internal/collection"
	"backend/internal/logger"
	collectionReq "backend/web/request/collection"
	"backend/web/response"
	"backend/web/util"
)

// AddCollection godoc
//
//	 @Title AddCollection
//		@Description	Get Subject by id
//	    @Security BearerAuth
//		@Tags			Collection
//		@Accept			json
//		@Produce		json
//	    @Param          collection 	body 	collectionReq.AddOrUpdateCollectionReq   true "collection"
//		@Success		200		 {object}   nil
//		@Router			/collection/:subject_id [post]
func (h Handler) AddCollection(c echo.Context) (err error) {
	var req collectionReq.AddOrUpdateCollectionReq
	var subjectId uint64
	param := c.Param("subject_id")
	if subjectId, err = strconv.ParseUint(param, 10, 32); err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		logger.Error("Failed to validate")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	// check
	if err := checkReq(req); err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	// get user id
	uid := c.Get("uid").(uint32)
	if err = h.ctrl.AddCollection(uid, uint32(subjectId), req); err != nil {
		logger.Error("Failed to add collection")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// UpdateCollection godoc
//
//	 @Title UpdateCollection
//		@Description	UpdateCollection
//	    @Security BearerAuth
//		@Tags			Collection
//		@Accept			json
//		@Produce		json
//	    @Param          collection 	body 	collectionReq.AddOrUpdateCollectionReq   true "collection"
//		@Success		200		 {object}   nil
//		@Router			/collection/:subject_id [patch]
func (h Handler) UpdateCollection(c echo.Context) error {
	var req collectionReq.AddOrUpdateCollectionReq
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Failed to validate")
		return util.Success(c, http.StatusBadRequest, err.Error())
	}
	uid := c.Get("uid").(uint32)
	subjectId, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}

	// check
	if err := checkReq(req); err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	err = h.ctrl.UpdateCollection(uid, uint32(subjectId), req)
	if err != nil {
		logger.Error("update failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// DeleteCollection godoc
// @Title DeleteCollection
//	    @Security BearerAuth
//	@Description	Delete Collection by subject_id
//	@Tags			Collection
//	@Accept			json
//	@Produce		json
//	@Param          subject_id 	path 	uint32   true "subject_id"
//	@Success		200			{object}  nil
//	@Router			/collection/:subject_id [delete]

func (h Handler) DeleteCollection(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	subjectId, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}
	err = h.ctrl.DeleteCollection(uid, uint32(subjectId))
	if err != nil {
		logger.Error("Failed to delete collection")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// GetMemberComment godoc
// @Title GetMemberComment
//
//	@Description	Get Comments By Uid
//	@Tags			Collection
//	@Accept			json
//	@Produce		json
//	@Param          member_id 	path 	uint32   true "member_id"
//	@Success		200			{object}   response.CommentsResp   "comments"
//	@Router			/:member_id/member/comment [get]
func (h Handler) GetMemberComment(c echo.Context) error {
	MemberId, err := strconv.ParseUint(c.Param("member_id"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}
	collections, err := h.ctrl.GetCollectionById("member", uint32(MemberId))
	if err != nil {
		logger.Error("get error")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	comments := util.Map(collections, response.NewCommentResp)
	responses := response.CommentsResp{
		Comments: comments,
	}
	return util.Success(c, http.StatusOK, responses)
}

// GetCommentsBySubjectID godoc
// @Title GetCommentsBySubjectID
//
//	@Description	Get Comments By Id
//	@Tags			Collection
//	@Accept			json
//	@Produce		json
//	@Param          subject_id 	path 	uint32   true "subject_id"
//	@Success		200			{object} response.CommentsResp  "comments"
//	@Router			/:subject_id/subject/comment [get]
func (h Handler) GetCommentsBySubjectID(c echo.Context) error {
	subjectId, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}
	collections, err := h.ctrl.GetCollectionById("subject", uint32(subjectId))
	if err != nil {
		logger.Error("get error")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	comments := util.Map(collections, response.NewCommentResp)
	responses := response.CommentsResp{
		Comments: comments,
	}
	return util.Success(c, http.StatusOK, responses)
}
func (h Handler) GetCollection(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	subjectId, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}
	Collection, err := h.ctrl.GetSelfCollection(uint32(subjectId), uid)
	if err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, Collection)
}

// GetSelfCollections godoc
// @Title GetSelfCollections
//
//	@Description	Get Comments By type
//	@Tags			Collection
//	@Accept			json
//	@Produce		json
//	@Param          type 	path 	uint8   true "type"
//	@Success		200			{object} ent.Collections  "collections"
//	@Router			/:type/collection [get]
func (h Handler) GetSelfCollections(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	Type, err := strconv.ParseUint(c.Param("type"), 10, 64)
	if err != nil {
		logger.Error("convert error")
		return util.Error(c, http.StatusNotFound, err.Error())
	}
	Collection, err := h.ctrl.GetSelfCollections(uid, uint8(Type))

	if err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, Collection)
}
func checkReq(req collectionReq.AddOrUpdateCollectionReq) error {
	if req.Type == collection.Wish && (req.Score != 0 || req.EpStatus != 0) {
		return errors.New("if type is wish, the score should be 0 or the ep_status should be 0")
	}
	if *req.HasComment == (req.Comment == "") {
		return errors.New("inconsistency between HasComment flag and Comment content")
	}
	return nil
}
