package blog

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/repository"
	"github.com/gogaruda/blog/blog/service"
)

type Module struct {
	TagService      service.TagService
	CategoryService service.CategoryService
}

func InitBlogModule(db *sql.DB) *Module {
	tagRepository := repository.NewTagRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	tagService := service.NewTagService(tagRepository)
	categoryService := service.NewCategoryService(categoryRepository)

	return &Module{
		TagService:      tagService,
		CategoryService: categoryService,
	}
}
