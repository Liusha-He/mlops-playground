package main

import (
	"fmt"
	"os"

	docs "users-service/docs"
	"users-service/src/middleware"
	"users-service/src/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     simple-service API
// @version         1.0
// @description     A Golang and gin API template
func main() {
	server := gin.New()
	docs.SwaggerInfo.BasePath = "/"

	server.Use(gin.Recovery(), middleware.Logger())

	// register all routes
	v1 := server.Group("/api/v1")

	routes.AuthRoutes(v1)
	routes.UserRoutes(v1)
	routes.HealthzRoutes(v1)

	// swagger ui documentation
	server.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(fmt.Sprintf(":%s", port))
	fmt.Println(fmt.Sprintf("Server running on port %s", port))
}
