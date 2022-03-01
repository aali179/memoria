package controllers

import (
	"memoria/app/models"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name string `json:"name"`
}

func CreateScrapbook(c *gin.Context) (models.Scrapbook, error) {
	body := CreateRequest{}
	if err := c.BindJSON(&body); err != nil {
		return models.Scrapbook{}, err
	} else {
		res, dberr := models.CreateScrapbook(body.Name)
		if dberr != nil {
			return models.Scrapbook{}, dberr
		}
		return res, nil

	}
}
