package servicesAcc

import (
	// "context"
	"fmt"
	"log"
	"os"
	"time"

	repositoriesAcc "github.com/DevOps-Group-D/YouToFy-API/repositories"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetAuthURL() string {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func GetPlaylist(username string, id string) ([]string, error) {
	token, err := GetYouTubeCredentials(username)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving YouTube credentials: %v", err)
	}
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %v", err)
	}
	ctx := context.Background()
	client := config.Client(ctx, token)
	opts := option.WithHTTPClient(client)
	service, err := youtube.NewService(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("Error creating YouTube service: %v", err)
	}
	part := []string{"snippet,contentDetails"}
	call := service.PlaylistItems.List(part).PlaylistId(id)
	var videoTitles []string
	err = call.Pages(ctx, func(response *youtube.PlaylistItemListResponse) error {
		for _, item := range response.Items {
			videoTitles = append(videoTitles, item.Snippet.Title)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error retrieving playlist items: %v", err)
	}
	return videoTitles, nil

}

func GetYouTubeCredentials(username string) (*oauth2.Token, error) {
	credentials, err := repositoriesAcc.GetYouTubeCredentials(username)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving YouTube credentials: %v", err)
	}
	var expiryTime time.Time
	if credentials.Expiry.Valid {
		expiryTime, err = time.Parse(time.RFC3339, credentials.Expiry.String)
		if err != nil {
			return nil, fmt.Errorf("Error parsing expiry time: %v", err)
		}
	}
	token := &oauth2.Token{
		AccessToken:  credentials.AccessToken,
		TokenType:    credentials.TokenType,
		RefreshToken: credentials.RefreshToken.String,
		Expiry:       expiryTime,
		ExpiresIn:    credentials.ExpiresIn.Int64,
	}
	return token, nil
}

func SaveToken(username string, token *oauth2.Token) error {
	err := repositoriesAcc.InsertYouTubeCredentials(
		username,
		token.AccessToken,
		token.TokenType,
		token.RefreshToken,
		token.Expiry.String(),
		token.ExpiresIn,
	)
	if err != nil {
		return fmt.Errorf("Error saving YouTube credentials: %v", err)
	}
	return nil
}

func GetWebTokenFromCode(code string) (*oauth2.Token, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %v", err)
	}
	ctx := context.Background()
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Unable to exchange code for token: %v", err)
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
