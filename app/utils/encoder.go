package utils

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func EncodeImage(file []byte) string {
	imageType := http.DetectContentType(file)
	fmt.Println(imageType)
	var decodedImage string
	if imageType == "image/jpeg" {
		decodedImage += "data:image/jpeg;base64,"
	} else {
		decodedImage += "data:image/png;base64,"
	}
	decodedImage += base64.StdEncoding.EncodeToString(file)

	return decodedImage
}
