package service

import (
	"github.com/gogaruda/blog/blog/model"
	"github.com/gogaruda/blog/blog/repository"
)

type TagService interface {
	GetAll() ([]model.TagModel, error)
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(r repository.TagRepository) TagService {
	return &tagService{repo: r}
}

func (s *tagService) GetAll() ([]model.TagModel, error) {
	return s.repo.GetAll()
}
