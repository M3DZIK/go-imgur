package imgur

type ImageInfoWithoutData struct {
	// Action successfully performed?
	Success bool `json:"success"`
	// Imgur HTTP Code
	Status int `json:"status"`
}

type ImageInfoData struct {
	// Image Info
	Data *ImageInfo `json:"data"`
	// Action successfully performed?
	Success bool `json:"success"`
	// Imgur HTTP Code
	Status int `json:"status"`
}

type ImageInfo struct {
	// Imgur img ID e.g. abcde123098
	ID string `json:"id"`
	// ID with extention e.g. abcde123098.png
	IDExt string
	// Image Title
	Title string `json:"title"`
	// Image Description
	Description string `json:"description"`
	// Image Upload Datetime
	Datetime int `json:"datetime"`
	// Type Format
	MimeType string `json:"type"`
	// Image is animated?
	Animated bool `json:"animated"`
	// Image width
	Width int `json:"width"`
	// Image height
	Height int `json:"height"`
	// Image File Size
	Size int `json:"size"`
	// Image Views
	Views int `json:"views"`
	// Imgur Bandwidth
	Bandwidth int `json:"bandwidth"`
	// Imgur Delete Hash
	Deletehash string `json:"deletehash,omitempty"`
	// The name of the file
	Name    string `json:"name,omitempty"`
	Section string `json:"section"`
	// Image URL e.g. https://i.imgur.com/abcde123098.png
	Link     string `json:"link"`
	Gifv     string `json:"gifv,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	Mp4Size  int    `json:"mp4_size,omitempty"`
	Looping  bool   `json:"looping,omitempty"`
	Favorite bool   `json:"favorite"`
	// Image is Nsfw?
	Nsfw      bool   `json:"nsfw"`
	Vote      string `json:"vote"`
	InGallery bool   `json:"in_gallery"`
}
