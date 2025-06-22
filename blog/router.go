package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gogaruda/auth/auth/middleware"
	"github.com/gogaruda/blog/blog/handler"
	"github.com/gogaruda/blog/blog/service"
	"github.com/gogaruda/pkg/validates"
)

func RegisterBlogRoutes(
	rg *gin.RouterGroup,
	tagService service.TagService,
) {
	v := validator.New()
	validation := validates.NewValidates(v)

	tagHandler := handler.NewTagHandler(tagService, validation)

	auth := rg.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/tags", middleware.RoleMiddleware(middleware.MatchAny, "admin", "editor", "penulis"), tagHandler.GetAllTags)
	auth.POST("/tags", middleware.RoleMiddleware(middleware.MatchAny, "admin", "editor"), tagHandler.CreateTag)
	auth.GET("/tags/:id", middleware.RoleMiddleware(middleware.MatchAny, "admin", "editor"), tagHandler.GetTagByID)
	auth.PUT("/tags/:id", middleware.RoleMiddleware(middleware.MatchAny, "admin", "editor"), tagHandler.UpdateTag)
	auth.DELETE("/tags/:id", middleware.RoleMiddleware(middleware.MatchAny, "admin", "editor"), tagHandler.DeleteTag)

}
