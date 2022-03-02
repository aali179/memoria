package controllers

import (
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
		res, err := GetPage(uint(id))
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
	res, err := CreatePage(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func createScrapbook(c *gin.Context) {
	res, err := CreateScrapbook(c)
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
	songs, err := SearchSong(song)
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
	song, err := GetSongById(songID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, song)
}

func getMap(c *gin.Context) {
	location := c.Query("location")
	googleMap, err := GetMapByLocation(location)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, googleMap)
}

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", getMessage)
	router.GET("/search", searchSongs)
	router.GET("/song", getSong)
	router.POST("/scrapbook", createScrapbook)
	router.POST("/page", createPage)
	router.GET("/page", getPage)
	router.GET("/map", getMap)

}
