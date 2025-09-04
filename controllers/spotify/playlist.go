package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	if len(urlParts) < 4 || urlParts[3] == "" {
		errMsg := "Error getting playlistId in URL"
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}
	playlistId := urlParts[3]

	playlist, err := p.Service.GetPlaylist(username, playlistId, accessToken.Value)
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
