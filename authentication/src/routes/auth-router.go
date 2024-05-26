package routes

import (
	"users-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/users/signup", controllers.Signup)
	r.POST("/users/login", controllers.Login)

}
