package user

import (
	"context"
	"net/http"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"backend/internal/logger"
	"backend/web/request/user"
	"backend/web/response"
	"backend/web/util"

	"github.com/labstack/echo/v4"
)

// Register POST
func (h Handler) Register(c echo.Context) error {
	// TODO Error handling
	var r user.UserRegisterReq
	if err := c.Bind(&r); err != nil {
		logger.Error("Register bind failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&r); err != nil {
		logger.Error("Register Validate failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	err := h.ctrl.Register(r)
	if err != nil {
		logger.Error("Failed to create new user")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) Login(c echo.Context) error {
	var req user.UserLoginReq
	var resp response.UserLoginResp
	if err := c.Bind(&req); err != nil {
		logger.Error("Login bind failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		logger.Error("Login Validate failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.ctrl.Login(req.Email, req.Password)
	if err != nil {
		logger.Error("Failed to login")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	resp.Token = token
	return util.Success(c, http.StatusOK, resp)
}

func (h Handler) Cancel(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	err := h.ctrl.Cancel(uid)
	if err != nil {
		logger.Error("Failed to cancel user")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) ModifyInfo(c echo.Context) (err error) {
	var req user.UserModifyInfoResp
	uid := c.Get("uid").(uint32)
	jsonData := c.FormValue("info")
	if err := jsoniter.Unmarshal([]byte(jsonData), &req.Info); err != nil {
		logger.Error("Failed to parse json")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	req.FileData, err = c.FormFile("image")
	if err != nil {
		logger.Error("Failed to get file")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	err = h.ctrl.ModifyInfo(uid, req)
	if err != nil {
		logger.Error("Failed to modify information")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) GetAvatar(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil || id <= 0 {
		logger.Error("Invalid id")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	url, err := h.ctrl.GetAvtar(uint32(id))
	if err != nil {
		logger.Error("Failed to get avtar")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, echo.Map{"url": url})
}

// GetMe godoc
//
//		 @Title GetMe
//			@Description	Get Me
//	     @Security BearerAuth
//			@Tags			Member
//			@Accept			json
//			@Produce		json
//			@Success		200		 {object}   response.UserGetResp   "MemberInfo"
//			@Router			/me [get]
func (h Handler) GetMe(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	member, err := h.userRepo.GetByUid(context.Background(), uid)
	if err != nil {
		logger.Error("Failed to get me")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, response.NewUserGetResp(member))
}
