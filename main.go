package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// single song retriever
// playlist retriever
type playlist struct {
	Songs []song
}

type song struct {
	Artists  []string `json:"artists"`
	Name     string   `json:"name"`
	Duration int      `json:"duration"`
}

func msToSeconds(duration int) int {
	trackSeconds := (duration / 1000)

	return trackSeconds
}

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

	// NOTE: requires playlist ID
	playlist, err := client.GetPlaylist("3CUKkyF0LGlg0HDSof03Cq")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	for _, track := range playlist.Tracks.Tracks {
		var artists []string
		for _, artist := range track.Track.Artists {
			artists = append(artists, artist.Name)
		}

		// calculate duration from int (ms) to 0:00 minute/seconds
		trackDurationSeconds := msToSeconds(track.Track.Duration)

		// remove 'original mix' and 'radio edit' from track name?

		// TODO: add date added to playlist + album?
		songPayload := &song{
			Artists: artists,
			Name:    track.Track.Name,
			// track duration in minutes
			Duration: trackDurationSeconds,
		}

		fmt.Println(songPayload)
	}

}
