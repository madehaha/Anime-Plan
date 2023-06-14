package subject

import (
	"backend/internal/logger"
	"backend/web/request"
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

	_ = c.JSON(http.StatusOK, subject)
	return nil
}
func (h Handler) CreateSubject(c echo.Context) error {
	var req request.CreateSubjectReq
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
	_ = c.JSON(http.StatusOK, "create success")
	return nil
}
