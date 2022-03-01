package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type Song struct {
	ID         spotify.ID `json:"ID"`
	Name       string     `json:"name"`
	ArtistName string     `json:"artist"`
	AlbumCover string     `json:"albumCover"`
	URL        string     `json:"url"`
}

var ctx = context.Background()

func connect() *spotify.Client {
	godotenv.Load()

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		fmt.Printf("couldn't get token: %s", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)

	client := spotify.New(httpClient)
	return client
}

func SearchSong(song string) []Song {
	client := connect()
	res, err := client.Search(ctx, song, spotify.SearchTypeTrack)

	if err != nil {
		fmt.Printf("something went wrong: %s", err)
	}
	songs := []Song{}

	if res.Tracks != nil {
		for _, track := range res.Tracks.Tracks {
			song := Song{Name: track.Name, ID: track.ID, ArtistName: track.Artists[0].Name}
			songs = append(songs, song)
		}
	}
	return songs
}

func GetSongById(ID string) Song {
	client := connect()
	track, err := client.GetTrack(ctx, spotify.ID(ID))

	if err != nil {
		fmt.Printf("something went wrong: %s", err)
	}
	song := Song{Name: track.Name, ID: track.ID, ArtistName: track.Artists[0].Name, URL: track.PreviewURL, AlbumCover: track.Album.Images[0].URL}
	return song

}
