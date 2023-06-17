package collection

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/internal/logger"
	collectionReq "backend/web/request/collection"
	"backend/web/util"
)

func (h Handler) AddCollection(c echo.Context) error {
	var req collectionReq.AddOrUpdateCollectionReq
	param := c.Param("subject_id")
	subjectId, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
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

	uid := c.Get("uid").(uint32)
	err = h.ctrl.AddCollection(uid, uint32(subjectId), req)
	if err != nil {
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
	err = h.ctrl.UpdateCollection(uid, uint32(subjectId), req)
	if err != nil {
		logger.Error("add failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// @Title DeleteCollection
//	@Description	Delete Collection by subject_id
//	@Tags			Collection
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}
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

func (h Handler) GetComment(c echo.Context) error {

	return nil
}
