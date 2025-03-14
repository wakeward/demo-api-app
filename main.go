package main

import (
	"github.com/wakeward/demo-api-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	controllers.Routes(router)

	router.Run(":8080")
}
