// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Test route
	router.GET("/test", getMessage)

	// Start serving the application
	router.Run()

}
