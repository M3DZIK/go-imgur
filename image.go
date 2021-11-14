package imgur

type ImageInfoWithoutData struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type ImageInfoData struct {
	Data    *ImageInfo `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

type ImageInfo struct {
	ID          string `json:"id"`
	IDExt       string
	Title       string `json:"title"`
	Description string `json:"description"`
	Datetime    int    `json:"datetime"`
	MimeType    string `json:"type"`
	Animated    bool   `json:"animated"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Size        int    `json:"size"`
	Views       int    `json:"views"`
	Bandwidth   int    `json:"bandwidth"`
	Deletehash  string `json:"deletehash,omitempty"`
	Name        string `json:"name,omitempty"`
	Section     string `json:"section"`
	Link        string `json:"link"`
	Gifv        string `json:"gifv,omitempty"`
	Mp4         string `json:"mp4,omitempty"`
	Mp4Size     int    `json:"mp4_size,omitempty"`
	Looping     bool   `json:"looping,omitempty"`
	Favorite    bool   `json:"favorite"`
	Nsfw        bool   `json:"nsfw"`
	Vote        string `json:"vote"`
	InGallery   bool   `json:"in_gallery"`
}
