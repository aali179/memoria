// routes.go

package main

import (
	"memoria/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
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

}
