package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	spotifyModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	authenticationService "github.com/DevOps-Group-D/YouToFy-API/services/authentication"
)

func (p spotifyProvider) Login(w http.ResponseWriter, r *http.Request) {
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

	authUrl := p.Service.GetAuthURL()

	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)

	response := "Spotify authorization route called"

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(302)
}

// TODO: Move it to authentication service
func (p spotifyProvider) Save(w http.ResponseWriter, r *http.Request) {
	var authReq spotifyModels.AuthenticationRequest

	err := json.NewDecoder(r.Body).Decode(&authReq)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding auth request body: %s", err.Error())
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

	authorized := authenticationService.Authorize(username.Value, r.Cookies())
	if !authorized {
		errMsg := "Error: Unauthorized"
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	accessToken, err := p.Service.GetAccessToken(username.Value, authReq.Code)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting access token %s", err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		fmt.Println(errMsg)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "spotify_access_token",
		Value:    accessToken.AccessToken,
		Expires:  time.Now().Add(time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(200)
}
