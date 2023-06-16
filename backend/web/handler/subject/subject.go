package subject

import (
	"backend/internal/logger"
	"backend/web/request/collection"
	"backend/web/request/subject"
	"backend/web/response"
	"backend/web/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h Handler) GetSubject(c echo.Context) error {
	subjects, err := h.ctrl.GetSubject()
	if err != nil {
		logger.Error("Failed to get subject")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	_ = c.JSON(http.StatusOK, subjects)
	return nil
}
func (h Handler) GetSubjectByID(c echo.Context) error {
	id := c.Param("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	subject, err := h.ctrl.GetSubjectByID(Id)
	if err != nil {
		logger.Error("Failed to find subject")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	return util.Success(c, http.StatusOK, subject)
}
func (h Handler) AddOrUpdateCollection(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	SubjectID := c.Param("id")
	AddType := c.Param("type")
	id, err := strconv.Atoi(SubjectID)
	if err != nil {
		logger.Error("add failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	subtype, err := strconv.Atoi(AddType)
	if err != nil {
		logger.Error("add failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	err = h.ctrl.AddCollection(uid, id, uint8(subtype))
	if err != nil {
		logger.Error("add failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
func (h Handler) GetComment(c echo.Context) error {
	return nil
}
func (h Handler) UpdateCollection(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	id := c.Param("id")
	var req collection.UpdateCollectionReq
	if err := c.Bind(&req); err != nil {
		logger.Error("Login bind failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Login Validate failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	subid, _ := strconv.Atoi(id)
	err := h.ctrl.UpdateCollection(uid, subid, req)
	if err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())

	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) SearchSubject(c echo.Context) error {
	searchTerm := c.QueryParam("term")
	if searchTerm == "" {
		logger.Error("fail to get term")
		return util.Error(c, http.StatusBadRequest, "term is nil")
	}
	Subject, err := h.ctrl.SearchSubject(searchTerm)
	if err != nil {
		logger.Error("fail to search")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, response.NewSubjectResp(Subject))
}

func (h Handler) CreateSubject(c echo.Context) error {
	var req subject.CreateSubjectReq
	if err := c.Bind(&req); err != nil {
		logger.Error("Login bind failed")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Login Validate failed")
		return err
	}
	gid := c.Get("gid").(uint8)
	err := h.ctrl.CreateSubject(req, gid)
	if err != nil {
		logger.Error("create subject failed")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	return c.NoContent(http.StatusOK)
}
