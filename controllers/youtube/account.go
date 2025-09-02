package youtube

import (
	"fmt"
	"net/http"

	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
)

func (p YoutubeProvider) Login(w http.ResponseWriter, r *http.Request) {
	authUrl := youtubeService.GetAuthURL()
	fmt.Println(authUrl)
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
	response := "Youtube authorization route called"

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(302)
	println("reached-end")
}

// func (p YoutubeProvider) save(w http.ResponseWriter, r *http.Request) {

// }
