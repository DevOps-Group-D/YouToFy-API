package servicesAcc

import (
	"fmt"
	"log"
	"os"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	youtubeRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/youtube"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	YOUTUBE_AUTH_URI                    = "https://accounts.google.com/o/oauth2/auth"
	YOUTUBE_TOKEN_URI                   = "https://oauth2.googleapis.com/token"
	YOUTUBE_AUTH_PROVIDER_X509_CERT_URL = "https://www.googleapis.com/oauth2/v1/certs"
	YOUTUBE_REDIRECT_URI                = "%s://%s"
)

// TODO: Add a struct to all these methods like Spotify
func GetAuthURL() string {
	config, err := loadFromConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	authURL := config.AuthCodeURL("youtube", oauth2.AccessTypeOffline)
	return authURL
}

func GetPlaylist(paylistId string, token *oauth2.Token) ([]string, error) {
	config, err := loadFromConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}
	ctx := context.Background()
	client := config.Client(ctx, token)
	opts := option.WithHTTPClient(client)
	service, err := youtube.NewService(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error creating YouTube service: %v", err)
	}
	part := []string{"snippet,contentDetails"}
	call := service.PlaylistItems.List(part).PlaylistId(paylistId)
	var videoTitles []string
	err = call.Pages(ctx, func(response *youtube.PlaylistItemListResponse) error {
		for _, item := range response.Items {
			videoTitles = append(videoTitles, item.Snippet.Title)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error retrieving playlist items: %v", err)
	}
	return videoTitles, nil

}

func GetYouTubeCredentials(username string) (*oauth2.Token, error) {
	credentials, err := youtubeRepository.GetYouTubeCredentials(username)
	if err != nil {
		return nil, fmt.Errorf("error retrieving YouTube credentials: %v", err)
	}
	token := &oauth2.Token{
		AccessToken: credentials.AccessToken,
	}
	return token, nil
}

func SaveToken(username string, token *oauth2.Token) error {
	err := youtubeRepository.InsertYouTubeCredentials(
		username,
		token.AccessToken,
	)
	if err != nil {
		return fmt.Errorf("error saving YouTube credentials: %v", err)
	}
	return nil
}

func GetWebTokenFromCode(code string) (*oauth2.Token, error) {
	config, err := loadFromConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}
	ctx := context.Background()
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("unable to exchange code for token: %v", err)
	}
	return token, nil
}

func loadConfig() (*oauth2.Config, error) {
	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config, nil
}

func loadFromConfig() (*oauth2.Config, error) {
	protocol := configs.Cfg.FrontConfig.Protocol
	host := configs.Cfg.FrontConfig.Host
	config := &oauth2.Config{
		ClientID:     configs.Cfg.YoutubeConfig.ClientId,
		ClientSecret: configs.Cfg.YoutubeConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  YOUTUBE_AUTH_URI,
			TokenURL: YOUTUBE_TOKEN_URI,
		},
		RedirectURL: fmt.Sprintf(YOUTUBE_REDIRECT_URI, protocol, host),
		Scopes:      []string{youtube.YoutubeReadonlyScope},
	}
	return config, nil
}
