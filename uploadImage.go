package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Upload Image to Imgur
//	info, status, err := client.UploadImage("https://example.com/example.png", "url", "abc")
func (client *Client) UploadImage(img string, dtype string, album string) (*ImageInfoData, int, error) {
	form := url.Values{}

	form.Add("image", img)
	form.Add("type", dtype)

	if album != "" {
		form.Add("album", album)
	}

	URL := "https://api.imgur.com/3/image"

	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, -1, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", client.Imgur.ClientID))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	if i.Status >= 300 {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	if !i.Success {
		return nil, i.Status, errors.New("Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	i.Data.IDExt = strings.Replace(i.Data.Link, "https://i.imgur.com/", "", -1)

	return &i, i.Status, nil
}

// Upload Image to Imgur by URL
//	info, status, err := client.UploadImageFromURL("https://abc/img", "")
func (client *Client) UploadImageFromURL(imgUrl string, album string) (*ImageInfoData, int, error) {
	return client.UploadImage(imgUrl, "url", album)
}

// Upload Image to Imgur by File
//	info, status, err := client.UploadImageFromURL("path/to/img", "")
func (client *Client) UploadImageFromFile(path string, album string) (*ImageInfoData, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, -1, err
	}
	defer f.Close()

	finfo, err := f.Stat()
	if err != nil {
		return nil, -1, err
	}

	size := finfo.Size()
	b := make([]byte, size)
	n, err := f.Read(b)
	if err != nil || int64(n) != size {
		return nil, -1, err
	}

	return client.UploadImage(string(b[:]), "file", album)
}
