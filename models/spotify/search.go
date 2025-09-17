package spotify

type FoundMusics struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Items []FoundItem `json:"items"`
}

type FoundItem struct {
	URI string `json:"uri"`
}
