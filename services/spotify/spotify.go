package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	spotifyModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	spotifyRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/spotify"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

type SpotifyService struct {
	Repository *spotifyRepository.SpotifyRepository
}

const (
	ACCESS_TOKEN_URL  = "https://accounts.spotify.com/api/token"
	AUTH_URL          = "https://accounts.spotify.com/authorize"
	GET_PLAYLISTS_URL = "https://api.spotify.com/v1/playlists/%s/tracks"
	REDIRECT_ROUTE    = "%s://%s"
)

var scopes = []string{"playlist-read-private", "playlist-modify-private", "playlist-modify-public"}

func (s *SpotifyService) GetAuthURL() string {
	protocol := configs.Cfg.FrontConfig.Protocol
	host := configs.Cfg.FrontConfig.Host
	redirectRoute := fmt.Sprintf(REDIRECT_ROUTE, protocol, host)

	state := utils.GenerateRandomString(16)
	formatedScopes := strings.Join(scopes, "%20")

	queryString := fmt.Sprintf("client_id=%s&response_type=%s&redirect_uri=%s&state=%s&scope=%s",
		configs.Cfg.SpotifyConfig.ClientId,
		"code",
		redirectRoute,
		state,
		formatedScopes,
	)

	authUrlWithParams := fmt.Sprintf("%s?%s", AUTH_URL, queryString)

	return authUrlWithParams
}

func (s *SpotifyService) GetAccessToken(username string, code string) (*spotifyModels.AccessTokenResponse, error) {
	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("redirect_uri", fmt.Sprintf("%s://%s", configs.Cfg.FrontConfig.Protocol, configs.Cfg.FrontConfig.Host))

	req, err := http.NewRequest(http.MethodPost, ACCESS_TOKEN_URL, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	authorization := fmt.Sprintf("%s:%s", configs.Cfg.SpotifyConfig.ClientId, configs.Cfg.SpotifyConfig.ClientSecret)

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(authorization)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := utils.Client.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var accessTokenRes spotifyModels.AccessTokenResponse

	err = json.Unmarshal(raw, &accessTokenRes)
	if err != nil {
		return nil, err
	}

	err = s.Repository.UpdateAccessToken(username, accessTokenRes.AccessToken)
	if err != nil {
		return nil, err
	}

	return &accessTokenRes, nil
}

func (s *SpotifyService) GetPlaylist(username string, playlistId string, accessToken string) (*spotifyModels.Playlist, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(GET_PLAYLISTS_URL, playlistId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := utils.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		errorBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Error making get spotify playlist request: %s", errorBody)
	}

	var playlist spotifyModels.Playlist

	err = json.NewDecoder(res.Body).Decode(&playlist)
	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
