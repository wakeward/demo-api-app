package main

import (
	"github.com/wakeward/demo-api-app/controllers"

	"github.com/gin-gonic/gin"
)

// Testing Gobfuscate Private GitHub Action 6
func main() {

	router := gin.Default()

	controllers.Routes(router)

	router.Run(":8080")
}
