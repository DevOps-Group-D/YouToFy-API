package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
	"golang.org/x/oauth2"
)

func (p YoutubeProvider) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := strings.Split(r.URL.String(), "/")[2]
	var authCode *oauth2.Token
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}
	err = json.Unmarshal(bodyBytes, &authCode)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
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
