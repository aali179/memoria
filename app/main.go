// main.go
package main

import (
	"fmt"
	"memoria/app/models"
	"memoria/app/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	Setup().Run()
}

//Setup routers to db
func Setup() *gin.Engine {
	router := gin.Default()
	err := models.Connect()
	api := &controllers.APIEnv{
		DB: models.GetDB(),
	}

	err := models.Connect()

	if err != nil {
		fmt.Println(err)
	}
	// Initialize all routes
	initializeRoutes(router)

	// Initialize all routes
	controllers.InitializeRoutes(router, api)

	return router
}