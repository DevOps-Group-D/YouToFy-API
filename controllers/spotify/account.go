package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	spotifyModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	spotifyRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/spotify"
	authenticationService "github.com/DevOps-Group-D/YouToFy-API/services/authentication"
	spotifyService "github.com/DevOps-Group-D/YouToFy-API/services/spotify"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

const ACCESS_TOKEN_URL = "https://accounts.spotify.com/api/token"

func Login(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting username cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	authorized := authenticationService.Authorize(username.Value)
	if !authorized {
		errMsg := "Error: Unauthorized"
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	authUrl := spotifyService.GetAuthURL()

	fmt.Println(authUrl)
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)

	response := "Spotify authorization route called"

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(302)
}

func Save(w http.ResponseWriter, r *http.Request) {
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

	authorized := authenticationService.Authorize(username.Value)
	if !authorized {
		errMsg := "Error: Unauthorized"
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", authReq.Code)
	body.Set("redirect_uri", fmt.Sprintf("%s://%s:%s/%s", configs.Cfg.FrontConfig.Protocol, configs.Cfg.FrontConfig.Host, configs.Cfg.FrontConfig.Port, "callback"))

	req, err := http.NewRequest(http.MethodPost, ACCESS_TOKEN_URL, strings.NewReader(body.Encode()))
	if err != nil {
		errMsg := fmt.Sprintf("Error creating post request: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	authorization := fmt.Sprintf("%s:%s", configs.Cfg.SpotifyConfig.ClientId, configs.Cfg.SpotifyConfig.ClientSecret)

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(authorization)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := utils.Client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error making authorization request: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	var accessTokenRes spotifyModels.AccessTokenResponse

	err = json.NewDecoder(res.Body).Decode(&accessTokenRes)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding access token response body: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	err = spotifyRepository.UpdateAccessToken(username.Value, accessTokenRes.AccessToken)
	if err != nil {
		errMsg := fmt.Sprintf("Error updating spotify access token: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "spotify_access_token",
		Value:    accessTokenRes.AccessToken,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	})
}
