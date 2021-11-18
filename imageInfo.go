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
	"strings"
)

// Get Image Info from Imgur
//	info, status, err := client.GetImageInfo("abc")
func (client *Client) GetImageInfo(imageID string) (*ImageInfoData, int, error) {
	form := url.Values{}

	URL := fmt.Sprintf("https://api.imgur.com/3/image/%s", imageID)

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, err
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	var i ImageInfoData

	err = dec.Decode(&i)
	if err != nil {
		return nil, -1, err
	}

	if !i.Success {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	i.Data.IDExt = strings.Replace(i.Data.Link, "https://i.imgur.com/", "", -1)

	return &i, i.Status, nil
}
