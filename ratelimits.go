package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type RateLimitsStruct struct {
	Data    RateLimitsData `json:"data"`
	Success bool           `json:"success"`
	Status  int            `json:"status"`
}

type RateLimitsData struct {
	UserLimit       int `json:"UserLimit"`
	UserRemaining   int `json:"UserRemaining"`
	UserReset       int `json:"UserReset"`
	ClientLimit     int `json:"ClientLimit"`
	ClientRemaining int `json:"ClientRemaining"`
}

// Get Image Info from Imgur
//	info, status, err := client.RateLimits()
func (client *Client) RateLimits() (*RateLimitsStruct, int, error) {
	form := url.Values{}

	URL := "https://api.imgur.com/3/credits"

	req, err := http.NewRequest("GET", URL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, -1, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", client.Imgur.ClientID))

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New("Imgur Failed with Status: " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	var i RateLimitsStruct

	err = dec.Decode(&i)
	if err != nil {
		return nil, -1, err
	}

	if !i.Success {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	return &i, i.Status, nil
}
