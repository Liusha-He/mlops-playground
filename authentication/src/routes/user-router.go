package routes

import (
	"users-service/src/controllers"
	"users-service/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.Use(middleware.Authentication())
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:user_id", controllers.GetUser)
}
