// routes.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/", getMessage)
}
