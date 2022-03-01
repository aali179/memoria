// routes.go

package main

import (
	"memoria/app/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func getPage(c *gin.Context) {
	pageID := c.Query("id")
	id, err := strconv.ParseUint(pageID, 10, 32)
	if err != nil {
		c.JSON(500, "Something went wrong")
	}
	c.JSON(200, controllers.GetPage(uint(id)))
}

func createPage(c *gin.Context) {
	c.JSON(200, controllers.CreatePage(c))
}

func searchSongs(c *gin.Context) {
	song := c.Query("song")
	songs := controllers.SearchSong(song)
	c.JSON(200, songs)
}

func getSong(c *gin.Context) {
	songID := c.Query("id")
	song := controllers.GetSongById(songID)
	c.JSON(200, song)
}
func initializeRoutes(router *gin.Engine) {
	router.GET("/", getMessage)
	router.GET("/search", searchSongs)
	router.GET("/song", getSong)
	router.POST("/page", createPage)
	router.GET("/page", getPage)

}
