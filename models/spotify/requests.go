package spotify

type AuthenticationRequest struct {
	Code  string `json:"code"`
	Error string `json:"error"`
	State string `json:"state"`
}

type InsertPlaylistRequest struct {
	Uris     []string `json:"uris"`
	Position int      `json:"position"`
}
