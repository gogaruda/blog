package service

import (
	"github.com/gogaruda/blog/blog/dto/request"
	"github.com/gogaruda/blog/blog/model"
	"github.com/gogaruda/blog/blog/repository"
)

type TagService interface {
	GetAll(limit, offset int) ([]model.TagModel, int, error)
	Create(req request.TagRequest) error
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) TagService {
	return &tagService{repo: r}
}

func (s *tagService) GetAll(limit, offset int) ([]model.TagModel, int, error) {
	tags, total, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, total, err
	}

	return tags, total, nil
}

func (s *tagService) Create(req request.TagRequest) error {
	exists, err := s.repo.CreateSlug(req.Name)
	if err != nil {
		return err
	}

	if err := s.repo.Create(req, exists); err != nil {
		return err
	}

	return nil
}
