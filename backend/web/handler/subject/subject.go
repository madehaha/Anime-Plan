package subject

import (
	"backend/web/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/internal/logger"
	"backend/web/request/subject"
	"backend/web/util"
)

// GetSubject godoc
//
//	 @Title GetSubject
//		@Description	Get All Subject
//		@Tags			Subject
//		@Accept			json
//		@Produce		json
//		@Success		200		 {object}   []response.GetSubjectResp "SubjectsInfo"
//		@Router			/subject/get [get]
func (h Handler) GetSubject(c echo.Context) error {
	subjects, err := h.ctrl.GetSubject()
	if err != nil {
		logger.Error("Failed to get subject")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}

	_ = c.JSON(http.StatusOK, util.Map(subjects, response.NewSubjectResp))
	return nil
}

// GetSubjectByID godoc
//
//	 @Title GetSubjectByID
//		@Description	Get Subject by id
//		@Tags			Subject
//		@Accept			json
//		@Produce		json
//	    @Param          subject_id 	path 	uint32   true "subject_id"
//		@Success		200		 {object}   response.GetSubjectResp "SubjectInfo"
//		@Router			/subject/:subject_id [get]
func (h Handler) GetSubjectByID(c echo.Context) error {
	id := c.Param("subject_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	subjectEntity, err := h.ctrl.GetSubjectByID(uint32(Id))
	if err != nil {
		logger.Error("Failed to find subject")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, response.NewSubjectResp(subjectEntity))
}

// TODO Use WikiJWTAuth

// CreateSubject godoc
//
//	 @Title CreateSubject
//	     @Security BearerAuth
//		@Description	Create Subject
//		@Tags			Subject
//		@Accept			json
//		@Produce		json
//	    @Param          SubjectReq body subject.CreateSubjectReq true "SubjectCreate"
//		@Success		200		 {object}  nil
//		@Router			/subject/create [post]
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
	if err := h.ctrl.CreateSubject(req); err != nil {
		logger.Error("create subject failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
