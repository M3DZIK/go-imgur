package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

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

	req.Header.Add("Authorization", "Client-ID "+client.ImgurClientID)
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

	if !i.Success {
		return nil, i.Status, errors.New("Upload to Imgur Failed with Status: " + strconv.Itoa(i.Status))
	}

	i.Data.IDExt = strings.Replace(i.Data.Link, "https://i.imgur.com/", "", -1)

	return &i, i.Status, nil
}

func (client *Client) UploadImageFromURL(imgUrl string, album string) (*ImageInfoData, int, error) {
	return client.UploadImage(imgUrl, "url", album)
}

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
