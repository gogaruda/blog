package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/gogaruda/auth/auth/middleware"
	"github.com/gogaruda/blog/blog/handler"
	"github.com/gogaruda/blog/blog/service"
)

func RegisterBlogRoutes(
	rg *gin.RouterGroup,
	tagService service.TagService,
) {
	tagHandler := handler.NewTagHandler(tagService)

	auth := rg.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		admin := auth.Group("/")
		admin.Use(middleware.RoleMiddleware(middleware.MatchAny, "super-admin"))
		{
			admin.GET("/tags", tagHandler.GetAllTags)
		}
	}
}
