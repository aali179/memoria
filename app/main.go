// main.go
package main

import (
	"fmt"
	"memoria/app/models"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	err := models.Connect()

	if err != nil {
		fmt.Println(err)
	}
	// Initialize all routes
	initializeRoutes(router)

	// Start serving the application
	router.Run()

}
