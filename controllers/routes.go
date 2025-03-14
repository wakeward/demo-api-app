package controllers

import (
	docs "github.com/wakeward/demo-api-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Demo API App
// @version         1.0
// @description     Simple Demo API Application

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func Routes(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/healthcheck", HealthCheck)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
