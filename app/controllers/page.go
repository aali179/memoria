package controllers

import (
	"memoria/app/models"

	"github.com/gin-gonic/gin"
)

func CreatePage(c *gin.Context) (models.Page, error) {
	//create image
	/*form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	mapImage := models.CreateImage(files["map[]"])
	ioutil.ReadFile(files["map"].Filename)
	imageOne := models.CreateImage([]byte(form["imageOne"]))
	imageTwo := models.CreateImage([]byte(c.PostForm("imageTwo")))
	imageThree := models.CreateImage([]byte(c.PostForm("imageThree")))

	location := models.CreateMap(mapImage, form["location"])
	song := models.CreateSong(form["spotifyID"])
	models.CreatePage(imageOne, imageTwo, imageThree, form["headingOne"], form["headingTwo"], form["headingThree"], location, song)*/

	newMap := models.Map{Location: "Mexico", ImageID: 1}
	song := models.Song{SpotifyID: "41293819"}
	res, err := models.CreatePage(models.Image{File: []byte("testing")}, models.Image{File: []byte("testing")}, models.Image{File: []byte("testing")}, "Test", "Test 2", "Test 3", newMap, song, 1)
	if err != nil {
		return models.Page{}, err
	}
	return res, nil
}

func GetPage(id uint) (models.Page, error) {
	res, err := models.GetPageById(id)
	if err != nil {
		return models.Page{}, err
	}
	return res, nil
}
