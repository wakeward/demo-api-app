package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// HealthCheck godoc
// @Summary healthcheck
// @Schemes
// @Description healthcheck for service
// @Tags healthcheck
// @Accept text/plain
// @Produce text/plain
// @Success 200 {string} HealthCheck
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	// Healthcheck response

	// Send response back
	c.String(http.StatusOK, "service is running")
}
