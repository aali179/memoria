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
	} else {
		res, err := controllers.GetPage(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func createPage(c *gin.Context) {
	res, err := controllers.CreatePage(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func createScrapbook(c *gin.Context) {
	res, err := controllers.CreateScrapbook(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)

}

func searchSongs(c *gin.Context) {
	song := c.Query("song")
	songs, err := controllers.SearchSong(song)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func getSong(c *gin.Context) {
	songID := c.Query("id")
	song, err := controllers.GetSongById(songID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, song)
}
func initializeRoutes(router *gin.Engine) {
	router.GET("/", getMessage)
	router.GET("/search", searchSongs)
	router.GET("/song", getSong)
	router.POST("/scrapbook", createScrapbook)
	router.POST("/page", createPage)
	router.GET("/page", getPage)

}
