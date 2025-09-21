package spotify

import "github.com/DevOps-Group-D/YouToFy-API/models"

type FoundMusics struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Items []FoundItem `json:"items"`
}

type FoundItem struct {
	URI string `json:"uri"`
}

type SearchPlaylist struct {
	Href  string        `json:"href"`
	Items []models.Item `json:"items"`
}
