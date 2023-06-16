package subject

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/internal/logger"
	"backend/web/request/subject"
	"backend/web/util"
)

func (h Handler) GetSubject(c echo.Context) error {
	subjects, err := h.ctrl.GetSubject()
	if err != nil {
		logger.Error("Failed to get subject")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	_ = c.JSON(http.StatusOK, subjects)
	return nil
}

func (h Handler) GetSubjectByID(c echo.Context) error {
	id := c.Param("subject_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	subjectEntity, err := h.ctrl.GetSubjectByID(uint32(Id))
	if err != nil {
		logger.Error("Failed to find subject")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, subjectEntity)
}

// TODO Use WikiJWTAuth
func (h Handler) CreateSubject(c echo.Context) error {
	var req subject.CreateSubjectReq
	if err := c.Bind(&req); err != nil {
		logger.Error("Create Subject request bind failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Create Subject request validate failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	gid := c.Get("gid").(uint8)
	err := h.ctrl.CreateSubject(req, gid)
	if err != nil {
		logger.Error("create subject failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
