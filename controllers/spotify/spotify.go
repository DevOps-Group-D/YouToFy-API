package spotify

import (
	"fmt"
	"net/http"

	spotifyService "github.com/DevOps-Group-D/YouToFy-API/services/spotify"
)

func Login(w http.ResponseWriter, r *http.Request) {
	authUrl := spotifyService.GetAuthURL()

	fmt.Println(authUrl)
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)

	response := "Spotify authorization route called"

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(200)
}
