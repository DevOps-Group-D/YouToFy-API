package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"

	youtubeModels "github.com/DevOps-Group-D/YouToFy-API/models/youtube"
	authenticationService "github.com/DevOps-Group-D/YouToFy-API/services/authentication"
	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
)

func (p youtubeProvider) Login(w http.ResponseWriter, r *http.Request) {
	authUrl := youtubeService.GetAuthURL()
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
	response := `{"authUrl": "` + authUrl + `"}`

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
	w.WriteHeader(302)
}

func (p youtubeProvider) Save(w http.ResponseWriter, r *http.Request) {
	var authReq youtubeModels.AuthenticationRequest

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

	authToken, error := youtubeService.GetWebTokenFromCode(authReq.Code)
	if error != nil {
		http.Error(w, "error retrieving token from code: "+error.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(authToken)
	if err != nil {
		http.Error(w, "error marshalling token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := youtubeService.SaveToken(username.Value, authToken); err != nil {
		http.Error(w, "error saving youtube token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "youtube_access_token",
		Value:    authToken.AccessToken,
		Expires:  authToken.Expiry,
		Path:     "/",
		HttpOnly: true,
	})

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}
