package spotify

import (
	"fmt"
	"strings"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

const (
	AUTH_URL       = "https://accounts.spotify.com/authorize"
	REDIRECT_ROUTE = "%s://%s:%s/callback"
)

var scopes = []string{"playlist-read-private", "playlist-modify-private", "playlist-modify-public"}

func GetAuthURL() string {
	protocol := configs.Cfg.FrontConfig.Protocol
	host := configs.Cfg.FrontConfig.Host
	port := configs.Cfg.FrontConfig.Port
	redirectRoute := fmt.Sprintf(REDIRECT_ROUTE, protocol, host, port)

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
