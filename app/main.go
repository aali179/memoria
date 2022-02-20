// main.go

package main

import (
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

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Initialize all routes
	initializeRoutes(router)

	// Start serving the application
	router.Run()

}
