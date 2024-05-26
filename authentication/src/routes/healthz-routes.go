package routes

import (
	"users-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func HealthzRoutes(r *gin.RouterGroup) {
	r.GET("heath/ready", controllers.Healthz)
}
