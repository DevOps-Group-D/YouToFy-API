package youtube

import (
	"encoding/json"
	"net/http"

	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
)

func (p YoutubeProvider) Login(w http.ResponseWriter, r *http.Request) {
	authUrl := youtubeService.GetAuthURL()
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
	response := `{"authUrl": "` + authUrl + `"}`

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(302)
}

func (p YoutubeProvider) Save(w http.ResponseWriter, r *http.Request) {
	// code := r.Header.Get("code")
	code := r.URL.Query().Get("code")
	authToken, error := youtubeService.GetWebTokenFromCode(code)
	if error != nil {
		http.Error(w, "error retrieving token from code: "+error.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(authToken)
	if err != nil {
		http.Error(w, "error marshalling token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}
