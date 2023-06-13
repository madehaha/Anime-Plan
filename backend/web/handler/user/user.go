package user

import (
	"backend/ent"
	"backend/internal/logger"
	"backend/web/request/user"

	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// Register POST
func (h Handler) Register(c echo.Context) error {
	// TODO Error handling
	var req user.UserRegisterReq
	if err := c.Bind(&req); err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	if err := c.Validate(&req); err != nil {
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
	_ = c.NoContent(http.StatusOK)
	return nil
}

func (h Handler) Login(c echo.Context) error {
	var req user.UserLoginReq
	if err := c.Bind(&req); err != nil {
		logger.Error("Login bind failed")
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Login Validate failed")
		return err
	}
	token, err := h.ctrl.Login(req.Email, req.Password)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}
	_ = c.JSON(http.StatusOK, token)
	logger.Info(token)
	return err
}

func (h Handler) Cancel(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	logger.Info("", zap.String("uid", strconv.Itoa(int(uid))))
	err := h.ctrl.Cancel(uid)
	if err != nil {
		logger.Error("Failed to cancel user")
		_ = c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	_ = c.NoContent(http.StatusOK)
	return nil
}
