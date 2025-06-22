package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/service"
	"github.com/gogaruda/pkg/apperror"
	"github.com/gogaruda/pkg/response"
	"github.com/gogaruda/pkg/validates"
)

type CategoryHandler struct {
	service  service.CategoryService
	Validate *validates.Validates
}

func NewCategoryHandler(s service.CategoryService, v *validates.Validates) *CategoryHandler {
	return &CategoryHandler{service: s, Validate: v}
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	fmt.Println(categories)

	response.OK(c, categories, "query ok", nil)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CategoryRequest
	if !h.Validate.ValidateJSON(c, &req) {
		return
	}

	if err := h.service.Create(req); err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	response.Created(c, nil, "query ok")
}
