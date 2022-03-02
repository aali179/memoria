package controllers

import (
	"context"
	"log"
	"fmt"

	"googlemaps.github.io/maps"
)

func GetMapByLocation(Location string) string {
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyCzLhQYVButqP7cEq1Xn0VcaghrGovpolE"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.StaticMapRequest{
		Center:   Location,
		Zoom:     14,
		Size:     "400x400",
		// Scale:    *scale,
		// Format:   maps.Format("png8"),
		// Language: *language,
		// Region:   *region,
		// MapType:  maps.MapType(*maptype),
	}

	resp, err := client.StaticMap(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	
}


