// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"memoria/app/models"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	models.Connect()
	// Initialize all routes
	initializeRoutes(router)

	// Start serving the application
	router.Run()

}
