package user

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/web/request"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Register POST
func (h Handler) Register(c echo.Context) error {
	// TODO
	var req request.UserRegisterReq
	if err := c.Bind(&req); err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	member := ent.Members{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		Nickname:     req.Nickname,
		RegisterTime: time.Now().Format("2006-01-02"),
	}

	err := h.ctrl.Register(&member)
	if err != nil {
		logger.Error("Failed to create new user")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	return nil
}
