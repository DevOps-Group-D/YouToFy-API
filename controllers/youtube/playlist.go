package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	authenticationService "github.com/DevOps-Group-D/YouToFy-API/services/authentication"
	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
	"golang.org/x/oauth2"
)

func (p youtubeProvider) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting username cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}
	authorized := authenticationService.Authorize(username.Value, r.Cookies())
	if !authorized {
		errMsg := "Error: Unauthorized"
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	playlistId := strings.Split(r.URL.String(), "/")[2]
	var authCode *oauth2.Token
	accessToken, err := r.Cookie("youtube_access_token")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting youtube_access_token cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}
	authCode = tokenFromHeader(accessToken.Value, "Bearer", accessToken.Expires)
	if authCode == nil {
		errMsg := "Error: Unauthorized"
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	playlist, err := youtubeService.GetPlaylist(playlistId, authCode)
	if err != nil {
		http.Error(w, "error retrieving playlist: "+err.Error(), http.StatusInternalServerError)
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
	w.Write([]byte(playlistJson))
	w.WriteHeader(http.StatusOK)
}

// TODO
func (p youtubeProvider) InsertPlaylist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Youtube] InsertPlaylist not implemented")
}

func tokenFromHeader(accessToken string, tokenType string, expiry time.Time) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   tokenType,
		Expiry:      expiry,
	}
}
