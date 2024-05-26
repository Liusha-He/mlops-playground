package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary ping
// @Schemes
// @Description do ping
// @Tags check
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @controller./api/v1/ping [get]
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, "service ready...")
}
