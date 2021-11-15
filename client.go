package imgur

import "net/http"

type Client struct {
	HTTPClient    *http.Client
	Imgur         Imgur
}

type Imgur struct {
	ClientID    string
	AccessToken string
}
