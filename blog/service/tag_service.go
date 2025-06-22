package service

import (
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/dto/response"
	"github.com/gogaruda/blog/blog/repository"
	"github.com/gogaruda/seo"
)

type TagService interface {
	GetAll(limit, offset int) ([]response.TagResponse, int, error)
	Create(req request.TagRequest) error
	GetByID(tagID string) (*response.TagResponse, error)
	Update(tagID string, request request.TagRequest) error
	Delete(tagID string) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) TagService {
	return &tagService{repo: r}
}

func (s *tagService) GetAll(limit, offset int) ([]response.TagResponse, int, error) {
	tags, total, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, total, err
	}

	return tags, total, nil
}

func (s *tagService) Create(req request.TagRequest) error {
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

func (s *tagService) GetByID(tagID string) (*response.TagResponse, error) {
	return s.repo.GetByID(tagID)
}

func (s *tagService) Update(tagID string, req request.TagRequest) error {
	tag, err := s.GetByID(tagID)
	if err != nil {
		return err
	}

	var seoDescription string
	if req.SeoDescription == "" {
		seoDescription = seo.Description(req.Description, 160)
	} else {
		seoDescription = seo.Description(req.SeoDescription, 160)
	}

	slug, err := s.repo.UpdateSlug(req.Name, tag.ID)
	if err != nil {
		return err
	}

	if err := s.repo.Update(req, slug, tag.ID, seoDescription); err != nil {
		return err
	}

	return nil
}

func (s *tagService) Delete(tagID string) error {
	tag, err := s.repo.GetByID(tagID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(tag.ID); err != nil {
		return err
	}

	return nil
}
