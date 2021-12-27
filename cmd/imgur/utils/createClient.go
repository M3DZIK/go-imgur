package utils

import (
	"net/http"

	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-imgur/cmd/imgur/config"
)

func CreateClient() imgur.Client {
	config := config.ParseConfig()

	return imgur.Client{
		HTTPClient: new(http.Client),
		Imgur: imgur.Imgur{
			ClientID: config.Imgur.ID,
		},
	}
}
