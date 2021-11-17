package test

import (
	"net/http"

	"github.com/MedzikUser/go-imgur"
)

func createClient(httpClient *http.Client, imgurClientID string, imgurClientSecret string) imgur.Client {
	var client = imgur.Client{
		HTTPClient: httpClient,
		Imgur: imgur.Imgur{
			ClientID:     imgurClientID,
			ClientSecret: imgurClientSecret,
		},
	}

	return client
}
