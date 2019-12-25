package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	playlist, err := client.GetPlaylist("3CUKkyF0LGlg0HDSof03Cq")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	for _, track := range playlist.Tracks.Tracks {
		fmt.Println("")
		for _, artists := range track.Track.Artists {
			fmt.Println(artists.Name)
		}
		fmt.Println(track.Track.Duration)
		fmt.Println(track.Track.Name)
	}

}
