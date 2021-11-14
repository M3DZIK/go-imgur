package imgur

import "net/http"

type Client struct {
	HTTPClient    *http.Client
	ImgurClientID string
}
