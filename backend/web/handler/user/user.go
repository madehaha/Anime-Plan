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

// Register godoc
//
//		 @Title Register
//			@Description	Register
//			@Tags			Member
//	        @Accept			json
//			@Produce		json
//			@Param  	registerReq 	body 	user.UserRegisterReq 	true "register req"
//			@Success		200		 {object}   nil " register success "
//			@Router			/register [post]
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

// Login godoc
//
//		 @Title Login
//			@Description	Login
//			@Tags			Member
//	        @Accept			json
//			@Produce		json
//			@Param  	loginReq 	body 	user.UserLoginReq 	true "login req"
//			@Success		200		 {object}   response.UserLoginResp " login success "
//			@Router			/login [post]
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

// Cancel godoc
//
//		 @Title Cancel
//			@Description	Cancel
//	     @Security BearerAuth
//			@Tags			Member
//	        @Accept			json
//			@Produce		json
//			@Success		200		 {object}   nil
//			@Router			/cancel [post]
func (h Handler) Cancel(c echo.Context) error {
	uid := c.Get("uid").(uint32)
	err := h.ctrl.Cancel(uid)
	if err != nil {
		logger.Error("Failed to cancel user")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// ModifyInfo godoc
//
//		 @Title ModifyInfo
//			@Description	ModifyInfo
//	     @Security BearerAuth
//			@Tags			Member
//	        @Accept			mpfd
//			@Produce		json
//			@Param         ModifyReq  body  user.ModifyInfoReq true "Modify Info"
//			@Success		200		 {object}   nil
//			@Router			/modify [put]
func (h Handler) ModifyInfo(c echo.Context) (err error) {
	var req user.ModifyInfoReq
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

// GetAvatar godoc
//
//			 @Title GetAvatar
//				@Description	Get Avatar
//		        @Security       BearerAuth
//				@Tags			Member
//				@Accept			json
//	         @Param          id      path   uint32   true "id"
//				@Produce		json
//				@Success		200		 {object}   string   "MemberAvatar"
//				@Router			/user/:id/avatar [get]
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

// GetMember godoc
//
//		 @Title GetMember
//			@Description	GetMember
//	     @Security BearerAuth
//			@Tags			Member
//			@Accept			json
//			@Produce		json
//			@Param       id    path   uint32    true  "member_id"
//			@Success		200		 {object}   response.UserGetResp   "MemberInfo"
//			@Router			/user/:member_id [get]
func (h Handler) GetMember(c echo.Context) error {
	id := c.Param("member_id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("error get id")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	member, err := h.userRepo.GetByUid(context.Background(), uint32(uid))
	if err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, response.NewUserGetResp(member))
}
