// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scrapbook struct {
	Page Page   `json:"page"`
	Name string `json:"name"`
}
type Page struct {
	ID           int    `json:"id"`
	Number       string `json:"number"`
	ImageOne     Image  `json:"imageOne"`
	ImageTwo     Image  `json:"imageTwo"`
	ImageThree   Image  `json:"imageThree"`
	HeadingOne   string `json:"headingOne"`
	HeadingTwo   string `json:"headingTwo"`
	HeadingThree string `json:"headingThree"`
	Song         Song   `json:"song"`
}

type Song struct {
	ID        int
	SpotifyID string
}

type Map struct {
	Location string
	Image    Image
}

type Image struct {
	ID int
}

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
