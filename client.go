package imgur

import "net/http"

type Client struct {
	HTTPClient *http.Client
	Imgur      Imgur
}

type Imgur struct {
	// Imgur Client-ID
	ClientID string
	// Imgur Bearer Token
	AccessToken string
}
