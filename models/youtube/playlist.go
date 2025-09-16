package youtube

import "time"

type Playlist struct {
	PlaylistID string `json:"playlist_id"`
	Items      []Item `json:"items,omitempty"`
	Uri        string `json:"uri"`
}

type Item struct {
	AddedAt time.Time `json:"added_at"`
	Track   Track     `json:"track"`
}

type Track struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Artists    []Artist `json:"artists"`
	Album      Album    `json:"album"`
	DurationMs int      `json:"duration_ms"`
	URI        string   `json:"uri"`
}

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Tracks struct {
	Items []Track `json:"items"`
}
