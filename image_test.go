package imgur

import (
	"net/http"
	"os"
	"testing"
)

var imgsDeleteHash []string

//* Upload

func TestUploadImageFromURL(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	i, _, err := client.UploadImageFromURL("https://golang.org/doc/gopher/fiveyears.jpg", "")
	if err != nil {
		t.Error("UploadImageFromFile() Failed with Error:", err)
		t.FailNow()
	}

	imgsDeleteHash = append(imgsDeleteHash, i.Data.Deletehash)
}

func TestUploadImageFromFile(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	i, _, err := client.UploadImageFromFile("test_data/fiveyears.jpg", "")
	if err != nil {
		t.Error("UploadImageFromFile() Failed with Error:", err)
		t.FailNow()
	}

	imgsDeleteHash = append(imgsDeleteHash, i.Data.Deletehash)
}

//* Delete

func TestDeleteImages(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	for _, v := range imgsDeleteHash {
		_, _, err := client.DeleteImageUnAuthed(v)
		if err != nil {
			t.Error("DeleteImageUnAuthed() Failed with Error:", err)
			t.FailNow()
		}
	}
}
