package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogaruda/blog/blog/service"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/response"
)

type TagHandler struct {
	service service.TagService
}

func NewTagHandler(s service.TagService) *TagHandler {
	return &TagHandler{service: s}
}

// GetAllTags godoc
// @Summary Get all users with pagination
// @Tags Tags
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.SwaggerResponse
// @Failure 401 {object} response.SwaggerResponse
// @Router /api/blog/tags [get]
func (h *TagHandler) GetAllTags(c *gin.Context) {
	tags, err := h.service.GetAll()
	if err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.OK(c, tags, "query ok", nil)
}
