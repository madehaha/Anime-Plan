package collection

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/internal/collection"
	"backend/internal/logger"
	collectionReq "backend/web/request/collection"
	"backend/web/util"
)

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
		return util.Success(c, http.StatusBadRequest, err.Error())
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

func (h Handler) GetComment(c echo.Context) error {
	return nil
}

func checkReq(req collectionReq.AddOrUpdateCollectionReq) error {
	if req.Type == collection.Wish && (req.Score != 0 || req.EpStatus != 0) {
		return errors.New("if type is wish, the score should be 0 or the ep_status should be 0")
	}
	if req.HasComment == (req.Comment == "") {
		return errors.New("inconsistency between HasComment flag and Comment content")
	}
	return nil
}
