package imgur

import "net/http"

// Imgur API Client
type Client struct {
	HTTPClient *http.Client
	Imgur      Imgur
}

type Imgur struct {
	// Imgur Client-ID
	ClientID string
	// Imgur Client Secret
	ClientSecret string

	// Imgur Access Token
	//	client.GenerateAccessToken()
	AccessToken string
}
