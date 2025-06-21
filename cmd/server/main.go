package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	authModule "github.com/gogaruda/auth/auth"
	"github.com/gogaruda/auth/auth/config"
	blogModule "github.com/gogaruda/blog/blog"
	"github.com/gogaruda/pkg/middleware"
	"os"
	"strings"
)

func getAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:3000"}
	}
	return strings.Split(origins, ",")
}

func main() {
	config.LoadENV()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := config.ConnectDB()
	auth := authModule.InitAuthModule(db)
	blog := blogModule.InitBlogModule(db)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware(getAllowedOrigins()))

	api := r.Group("/api")

	// Module Auth
	authModule.RegisterAuthRoutes(api.Group("/auth"), auth.AuthService, auth.UserService)

	// Module Blog
	blogModule.RegisterBlogRoutes(api.Group("/blog"), blog.TagService)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	_ = r.Run(":" + port)
}
