package imgur

import (
	"net/http"
)

func createClient(httpClient *http.Client, imgurClientID string) Client {
	var client = Client{
		HTTPClient: httpClient,
		Imgur: Imgur{
			ClientID: imgurClientID,
		},
	}

	return client
}
