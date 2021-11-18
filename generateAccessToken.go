package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AccessTokenResponse struct {
	// Access Token
	AccessToken string `json:"access_token"`
	// Token Expires Time
	Expires uint `json:"expires_in"`
	// Tonen Type "bearer"
	TokenType string `json:"token_type"`
	// Refresh Token
	RefreshToken string `json:"refresh_token"`
	// Imgur Account ID
	AccountID uint `json:"account_id"`
	// Account Username
	Username string `json:"account_username"`
}

// Generate Access Token (Bearer)
//	res, err := client.GenerateAccessToken("abc")
func (client *Client) GenerateAccessToken(refreshToken string) (*AccessTokenResponse, error) {
	form := url.Values{}

	form.Add("refresh_token", refreshToken)
	form.Add("client_id", client.Imgur.ClientID)
	form.Add("client_secret", client.Imgur.ClientSecret)
	form.Add("grant_type", "refresh_token")

	URL := "https://api.imgur.com/oauth2/token"

	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 && res.StatusCode >= 300 {
		return nil, errors.New("Imgur Failed with code: " + strconv.Itoa(res.StatusCode))
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	var i AccessTokenResponse

	err = dec.Decode(&i)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
