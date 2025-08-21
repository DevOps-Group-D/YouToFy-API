package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	spotifyModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	"github.com/DevOps-Group-D/YouToFy-API/services/authentication"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

const GET_PLAYLISTS_URL = "https://api.spotify.com/v1/playlists/%s/tracks"

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := strings.Split(r.URL.String(), "/")[3]

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(GET_PLAYLISTS_URL, playlistId), nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating get spotify playlist request: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

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

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.Value))

	res, err := utils.Client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error making get spotify playlist request: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	if res.StatusCode != http.StatusOK {
		errorBody, err := io.ReadAll(res.Body)
		if err != nil {
			errMsg := fmt.Sprintf("Error reading error body: %s", errorBody)
			http.Error(w, errMsg, http.StatusBadRequest)
			fmt.Println(errMsg)
			return
		}

		errMsg := fmt.Sprintf("Error making get spotify playlist request: %s", errorBody)
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	var playlistt spotifyModels.Playlist

	err = json.NewDecoder(res.Body).Decode(&playlistt)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding playlist: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	a, _ := json.Marshal(playlistt)
	fmt.Println(string(a))
}
