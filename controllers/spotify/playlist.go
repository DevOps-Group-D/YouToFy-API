package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	spotifyModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	"github.com/DevOps-Group-D/YouToFy-API/services/authentication"
)

func (p spotifyProvider) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		errMsg := fmt.Sprintf("Missing username cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}
	username := c.Value

	authorized := authentication.Authorize(username, r.Cookies())
	if !authorized {
		errMsg := "Error: unauthorized"
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	accessToken, err := r.Cookie("spotify_access_token")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting spotify_access_token cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	urlParts := strings.Split(r.URL.String(), "/")
	if len(urlParts) < 3 || urlParts[2] == "" {
		errMsg := "Error getting playlistId in URL"
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}
	playlistId := urlParts[2]

	playlist, err := p.Service.GetPlaylist(playlistId, accessToken.Value)
	if err != nil {
		errMsg := "Error getting playlist"
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	playlistJson, err := json.Marshal(playlist)
	if err != nil {
		errMsg := "Error marshalling playlist"
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(playlistJson))
}

func (p spotifyProvider) InsertPlaylist(w http.ResponseWriter, r *http.Request) {
	var playlist *spotifyModels.Playlist

	err := json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding playlist request body: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	username, err := r.Cookie("username")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting username cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	authorized := authentication.Authorize(username.Value, r.Cookies())
	if !authorized {
		errMsg := "Error: unauthorized"
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	accessToken, err := r.Cookie("spotify_access_token")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting spotify_access_token cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	urlParts := strings.Split(r.URL.String(), "/")
	if len(urlParts) < 3 || urlParts[2] == "" {
		errMsg := "Error getting playlistId in URL"
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}
	playlistId := urlParts[2]

	err = p.Service.InsertPlaylist(playlistId, username.Value, accessToken.Value, playlist)
	if err != nil {
		errMsg := fmt.Sprintf("Error inserting musics into playlist: %s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	playlist, err = p.Service.GetPlaylist(playlistId, accessToken.Value)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting playlist: %s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	playlistJson, err := json.Marshal(playlist)
	if err != nil {
		errMsg := fmt.Sprintf("Error marshalling playlist: %s", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(playlistJson))
	w.WriteHeader(200)
}
