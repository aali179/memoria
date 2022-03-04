package controllers

import (
	"bytes"
	"context"
	"image/png"
	"memoria/app/utils"
	"os"

	"googlemaps.github.io/maps"
)

type Map struct {
	Location string `json:"location"`
	Image    string `json:"image"`
}

// Setup maps client
func connectMaps() (*maps.Client, error) {
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("MAPS_KEY")))
	if err != nil {
		return nil, err
	}
	return client, err
}

func GetMapByLocation(Location string) (Map, error) {
	client, err := connectMaps()

	if err != nil {
		return Map{}, err
	}

	// Constructing marker
	marker := maps.Marker{Color: "#E26D5A", Size: "mid", LocationAddress: Location}

	r := &maps.StaticMapRequest{
		Center:  Location,
		Zoom:    14,
		Size:    "400x400", //change to page dimensions
		Markers: []maps.Marker{marker},
	}

	resp, err := client.StaticMap(context.Background(), r)

	if err != nil {
		return Map{}, err
	}

	//Encoding from Image to base64
	//In-memory buffer to store PNG image before we base 64 encode it
	var buff bytes.Buffer
	png.Encode(&buff, resp)
	imageEncoded := utils.EncodeImage(buff.Bytes())

	if err != nil {
		return Map{}, err
	}

	googleMap := Map{Location: Location, Image: imageEncoded}

	return googleMap, nil

}
