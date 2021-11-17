package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/MedzikUser/go-imgur"
)

func generateAccessToken() (imgur.Client, error) {
	id := os.Getenv("IMGUR_CLIENT_ID")
	secret := os.Getenv("IMGUR_CLIENT_SECRET")
	refresh := os.Getenv("IMGUR_REFRESH_TOKEN")

	client := createClient(new(http.Client), id, secret)
	res, err := client.GenerateAccessToken(refresh)

	client.Imgur.AccessToken = res.AccessToken

	return client, err
}

// Upload image from URL
func TestGenerateAccessToken(t *testing.T) {
	id := os.Getenv("IMGUR_CLIENT_ID")
	secret := os.Getenv("IMGUR_CLIENT_SECRET")
	refresh := os.Getenv("IMGUR_REFRESH_TOKEN")

	client := createClient(new(http.Client), id, secret)

	_, err := client.GenerateAccessToken(refresh)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
