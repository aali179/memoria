// main.go
package main

import (
	"fmt"
	"memoria/app/controllers"
	"memoria/app/models"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	// router := gin.Default()

	// err := models.Connect()
	// api := &controllers.APIEnv{
	// 	DB: models.GetDB(),
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // Initialize all routes
	// controllers.InitializeRoutes(router, api)
	Setup().Run()
	SetupTest().Run()

	// Start serving the application
	//router.Run()

}

//Setup routers to db
func Setup() *gin.Engine {
	router := gin.Default()
	err := models.Connect()
	api := &controllers.APIEnv{
		DB: models.GetDB(),
	}

	if err != nil {
		fmt.Println(err)
	}

	// Initialize all routes
	controllers.InitializeRoutes(router, api)

	return router
}

//Setup routers to test db
func SetupTest() *gin.Engine {
	router := gin.Default()
	err := models.ConnectTestDB()
	api := &controllers.APIEnv{
		DB: models.GetTestDB(),
	}

	if err != nil {
		fmt.Println(err)
	}

	// Initialize all routes
	controllers.InitializeRoutes(router, api)

	return router
}
