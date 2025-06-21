package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	authModule "github.com/gogaruda/auth/auth"
	"github.com/gogaruda/auth/auth/config"
	_ "github.com/gogaruda/auth/docs"
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

// Swagger documentation
// @title Blog - REST API Docs
// @description Blog system
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
