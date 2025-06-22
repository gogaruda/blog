package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/service"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/response"
	"github.com/gogaruda/pkg/validates"
	"strconv"
)

type TagHandler struct {
	service   service.TagService
	Validator *validates.Validates
}

func NewTagHandler(s service.TagService, v *validates.Validates) *TagHandler {
	return &TagHandler{service: s, Validator: v}
}

func (h *TagHandler) GetAllTags(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	tags, total, err := h.service.GetAll(limit, offset)
	if err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	meta := response.MetaData{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	response.OK(c, tags, "query ok", &meta)
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var req request.TagRequest
	if !h.Validator.ValidateJSON(c, &req) {
		return
	}

	if err := h.service.Create(req); err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.Created(c, nil, "query ok")
}

func (h *TagHandler) GetTagByID(c *gin.Context) {
	tagID := c.Param("id")
	tag, err := h.service.GetByID(tagID)
	if err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.OK(c, tag, "query ok", nil)
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	var req request.TagRequest
	if !h.Validator.ValidateJSON(c, &req) {
		return
	}

	if err := h.service.Update(c.Param("id"), req); err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.Created(c, nil, "query ok")
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.NoContent(c)
}
