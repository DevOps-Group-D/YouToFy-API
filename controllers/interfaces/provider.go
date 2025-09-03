package interfaces

import "net/http"

type Provider interface {
	Login(w http.ResponseWriter, r *http.Request)
	// Save(w http.ResponseWriter, r *http.Request)
	GetPlaylist(w http.ResponseWriter, r *http.Request)
}
