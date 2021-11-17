package imgur

import (
	"net/http"
	"os"
	"testing"
)

var imgDeleteHashs []string

// Upload image from URL
func TestUploadImageFromURL(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	i, _, err := client.UploadImageFromURL("https://golang.org/doc/gopher/fiveyears.jpg", "")
	if err != nil {
		t.Error("UploadImageFromFile() Failed with Error:", err)
		t.FailNow()
	}

	imgDeleteHashs = append(imgDeleteHashs, i.Data.Deletehash)
}

// Upload image from File
func TestUploadImageFromFile(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	i, _, err := client.UploadImageFromFile("test_data/fiveyears.jpg", "")
	if err != nil {
		t.Error("UploadImageFromFile() Failed with Error:", err)
		t.FailNow()
	}

	imgDeleteHashs = append(imgDeleteHashs, i.Data.Deletehash)
}

// Delete images
func TestDeleteImages(t *testing.T) {
	key := os.Getenv("IMGUR_CLIENT_ID")

	client := createClient(new(http.Client), key)

	for _, v := range imgDeleteHashs {
		_, _, err := client.DeleteImageUnAuthed(v)
		if err != nil {
			t.Error("DeleteImageUnAuthed() Failed with Error:", err)
			t.FailNow()
		}
	}
}
