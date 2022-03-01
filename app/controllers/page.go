package controllers

import (
	"io/ioutil"
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
	mapImage, err := ioutil.ReadFile("./helpers/images/image-asset.jpeg")
	if err != nil {
		return models.Page{}, err
	}
	createdImage, err := models.CreateImage(mapImage)

	if err != nil {
		return models.Page{}, err
	}
	newMap := models.Map{Location: "Mexico", ImageID: createdImage.ID}
	song := models.Song{SpotifyID: "41293819"}
	fileArray := []string{
		"./helpers/images/image-asset.jpeg",
		"./helpers/images/image-asset.jpeg",
	}
	images := []models.Image{}

	for _, path := range fileArray {
		imgFile, err := ioutil.ReadFile(path)
		if err != nil {
			return models.Page{}, err
		}
		image, err := models.CreateImage(imgFile)
		if err != nil {
			return models.Page{}, err
		}
		images = append(images, image)
	}

	res, err := models.CreatePage(images, "Test", "Test 2", "Test 3", newMap, song, 1)
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
