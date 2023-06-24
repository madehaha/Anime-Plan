package subject

import (
	"backend/ent"
	"backend/web/response"
	jsoniter "github.com/json-iterator/go"

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

	_ = c.JSON(http.StatusOK, util.Map(subjects, response.SubjectResp))
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
//		@Success		200		 {object}   response.GetSubjectWithFieldResp "SubjectInfo"
//		@Router			/subject/:subject_id [get]
func (h Handler) GetSubjectByID(c echo.Context) error {
	id := c.Param("subject_id")
	Id, err := strconv.ParseUint(id, 10, 64)
	subjectEntity, Field, err := h.ctrl.GetSubjectByID(uint32(Id))
	if err != nil {
		logger.Error("Failed to find subject")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	Subjectwithfield := response.SubjectWithField{
		Subject: subjectEntity,
		Field:   Field,
	}
	return util.Success(c, http.StatusOK, response.NewSubjectResp(Subjectwithfield))
}

// TODO Use WikiJWTAuth
func (h Handler) SearchSubject(c echo.Context) error {
	name := c.QueryParam("name")
	subjectEntity, Field, err := h.ctrl.GetSubjectByName(name)
	if err != nil {
		logger.Error("search failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	subjectwithfield := response.SubjectWithField{
		Field:   Field,
		Subject: subjectEntity,
	}
	return util.Success(c, http.StatusOK, subjectwithfield)
}

func (h Handler) Rankings(c echo.Context) error {
	subjects, err := h.ctrl.Rankings()
	if err != nil {
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	length := len(subjects)
	res := make([]response.SubjectWithField, length)

	for key, value := range subjects {
		res[key].Subject = new(ent.Subject)
		res[key].Subject = value.Subject
		res[key].Field = new(ent.SubjectField)
		res[key].Field = value.Field
	}

	return util.Success(c, http.StatusOK, util.Map(res, response.NewSubjectResp))
}

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

// CreateSubjectWithSave godoc
//
//	 @Title CreateSubjectWithSave
//	     @Security BearerAuth
//		@Description	Create Subject
//		@Tags			Subject
//		@Accept			mpfd
//		@Produce		json
//	    @Param          SubjectReq body subject.CreateSubjectWithSaveReq true "SubjectCreate"
//		@Success		200		 {object}  nil
//		@Router			/subject/create [put]
func (h Handler) CreateSubjectWithSave(c echo.Context) (err error) {
	var req subject.CreateSubjectWithSaveReq
	uid := c.Get("uid").(uint32)
	gid := c.Get("gid").(uint8)
	if gid == 0 {
		return util.Error(c, http.StatusBadRequest, "no permission:have no root")
	}
	jsonData := c.FormValue("createSubject")
	if err := jsoniter.Unmarshal([]byte(jsonData), &req.CreateSubject); err != nil {
		logger.Error("Failed to parse json")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	req.FileData, err = c.FormFile("image")
	if err != nil {
		logger.Error("Failed to get file")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	if err := h.ctrl.CreateSubjectWithSave(uid, req); err != nil {
		logger.Error("create subject with save failed")
		return util.Error(c, http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) BoardCast(c echo.Context) error {
	id := c.Param("weekday")
	Day, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error("error day")
		return util.Success(c, http.StatusBadRequest, err.Error())
	}
	Board, err := h.ctrl.BoardCast(uint8(Day))
	if err != nil {
		logger.Error("get boardcast error")
		return util.Success(c, http.StatusBadRequest, err.Error())
	}
	return util.Success(c, http.StatusOK, util.Map(Board, response.SubjectResp))
}
