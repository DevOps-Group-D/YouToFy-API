package interfaces

import "net/http"

// Interacts with music providers e.g (YouTube, Spotify...)
type Provider interface {

	// Login on account
	Login(w http.ResponseWriter, r *http.Request)

	// Save access_token on application account
	Save(w http.ResponseWriter, r *http.Request)

	// Get musics from a specific playlist
	GetPlaylist(w http.ResponseWriter, r *http.Request)

	// Insert musics into a specific playlist
	InsertPlaylist(w http.ResponseWriter, r *http.Request)
}
