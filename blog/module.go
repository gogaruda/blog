package blog

import (
	"database/sql"
	"github.com/gogaruda/blog/blog/repository"
	"github.com/gogaruda/blog/blog/service"
)

type Module struct {
	TagService service.TagService
}

func InitBlogModule(db *sql.DB) *Module {
	tagRepository := repository.NewTagRepository(db)

	tagService := service.NewTagService(tagRepository)

	return &Module{
		TagService: tagService,
	}
}
