package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func (a *APIEnv) getPage(c *gin.Context) {
	pageID := c.Query("id")
	id, err := strconv.ParseUint(pageID, 10, 32)
	if err != nil {
		c.JSON(500, "Something went wrong")
	} else {
		res, err := a.GetPage(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (a *APIEnv) createPage(c *gin.Context) {
	res, err := a.CreatePage(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (a *APIEnv) createScrapbook(c *gin.Context) {
	res, err := a.CreateScrapbook(c)
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

func InitializeRoutes(router *gin.Engine, api *APIEnv) {
	router.GET("/", getMessage)
	router.GET("/search", searchSongs)
	router.GET("/song", getSong)
	router.POST("/scrapbook", api.createScrapbook)
	router.POST("/page", api.createPage)
	router.GET("/page", api.getPage)
	router.GET("/map", getMap)

}
