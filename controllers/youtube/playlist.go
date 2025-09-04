package youtube

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	authenticationService "github.com/DevOps-Group-D/YouToFy-API/services/authentication"
	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
	"golang.org/x/oauth2"
)

func (p YoutubeProvider) GetPlaylist(w http.ResponseWriter, r *http.Request) {
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
		errMsg := fmt.Sprintf("Error getting spotify_access_token cookie: %s", err.Error())
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

	videos, err := youtubeService.GetPlaylist(playlistId, authCode)
	if err != nil {
		http.Error(w, "error retrieving playlist: "+err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse := `{"videos": ["` + strings.Join(videos, `","`) + `"]}`
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
	w.WriteHeader(http.StatusOK)
}

func tokenFromHeader(accessToken string, tokenType string, expiry time.Time) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   tokenType,
		Expiry:      expiry,
	}
}
