package servicesAcc

import (
	"fmt"
	"log"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	youtubeModels "github.com/DevOps-Group-D/YouToFy-API/models/spotify"
	youtubeRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/youtube"

	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
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

func GetPlaylist(paylistId string, token *oauth2.Token) (youtubeModels.Playlist, error) {
	config, err := loadFromConfig()
	if err != nil {
		return youtubeModels.Playlist{}, fmt.Errorf("error loading config: %v", err)
	}
	ctx := context.Background()
	client := config.Client(ctx, token)
	opts := option.WithHTTPClient(client)
	service, err := youtube.NewService(ctx, opts)
	if err != nil {
		return youtubeModels.Playlist{}, fmt.Errorf("error creating YouTube service: %v", err)
	}
	part := []string{"snippet,contentDetails"}
	call := service.PlaylistItems.List(part).PlaylistId(paylistId)
	var playlist youtubeModels.Playlist
	playlist.PlaylistID = paylistId
	playlist.Uri = paylistId
	var items []youtubeModels.Item
	playlist.Items = items
	err = call.Pages(ctx, func(response *youtube.PlaylistItemListResponse) error {
		for _, item := range response.Items {
			artist := youtubeModels.Artist{
				ID:   item.Snippet.VideoOwnerChannelId,
				Name: item.Snippet.VideoOwnerChannelTitle,
			}
			var artists []youtubeModels.Artist
			artists = append(artists, artist)
			var images []youtubeModels.Image
			default_thumb := item.Snippet.Thumbnails.Default
			if default_thumb != nil {
				image := youtubeModels.Image{
					Height: int(default_thumb.Height),
					URL:    default_thumb.Url,
					Width:  int(default_thumb.Width),
				}
				images = append(images, image)
			}
			standard_thumb := item.Snippet.Thumbnails.Standard
			if standard_thumb != nil {
				image := youtubeModels.Image{
					Height: int(standard_thumb.Height),
					URL:    standard_thumb.Url,
					Width:  int(standard_thumb.Width),
				}
				images = append(images, image)
			}
			album := youtubeModels.Album{
				ID:     "youtube",
				Name:   "youtube",
				Images: images,
			}

			track := youtubeModels.Track{
				ID:         item.Id,
				Name:       item.Snippet.Title,
				Artists:    artists,
				Album:      album,
				DurationMs: 0,
				URI:        item.ContentDetails.VideoId,
			}
			publishedAtTime, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				publishedAtTime = time.Time{}
			}
			item := youtubeModels.Item{
				AddedAt: publishedAtTime,
				Track:   track,
			}
			playlist.Items = append(playlist.Items, item)
		}
		return nil
	})
	if err != nil {
		return youtubeModels.Playlist{}, fmt.Errorf("error retrieving playlist items: %v", err)
	}
	return playlist, nil

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
	_, Credentials_err := youtubeRepository.GetYouTubeCredentials(username)
	// fmt.Println(res)
	// fmt.Printf("will update credentials?: %v\n", Credentials_err)
	if Credentials_err == nil {
		// fmt.Println("updating credentials")
		err := youtubeRepository.UpdateYouTubeCredentials(username, token.AccessToken)
		if err != nil {
			return fmt.Errorf("error updating YouTube credentials: %v", err)
		}
	} else {
		err := youtubeRepository.InsertYouTubeCredentials(
			username,
			token.AccessToken,
		)
		if err != nil {
			return fmt.Errorf("error saving YouTube credentials: %v", err)
		}
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
