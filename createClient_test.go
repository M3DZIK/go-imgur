package imgur

import (
	"net/http"
)

func createClient(httpClient *http.Client, imgurClientID string) *Client {
	client := new(Client)
	client.HTTPClient = httpClient
	client.ImgurClientID = imgurClientID

	return client
}
