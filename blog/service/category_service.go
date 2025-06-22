package service

import (
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/dto/response"
	"github.com/gogaruda/blog/blog/repository"
	"github.com/gogaruda/seo"
)

type CategoryService interface {
	GetAll() ([]*response.CategoryResponse, error)
	Create(req request.CategoryRequest) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) GetAll() ([]*response.CategoryResponse, error) {
	return s.repo.GetAll()
}

func (s *categoryService) Create(req request.CategoryRequest) error {
	slug, err := s.repo.CreateSlug(req.Name)
	if err != nil {
		return err
	}

	var seoDescription string
	if req.SeoDescription == "" {
		seoDescription = seo.Description(req.Description, 160)
	} else {
		seoDescription = seo.Description(req.SeoDescription, 160)
	}

	if err := s.repo.Create(req, slug, seoDescription); err != nil {
		return err
	}

	return nil
}
