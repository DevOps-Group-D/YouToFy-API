package configs

import (
	"fmt"
	"os"

	"github.com/DevOps-Group-D/YouToFy-API/utils"
	"github.com/spf13/viper"
)

type config struct {
	AuthenticationConfig *AuthenticationConfig
	DBConfig             *DBConfig
	FrontConfig          *FrontConfig
	Provider             utils.Provider
	SpotifyConfig        *SpotifyConfig
	YoutubeConfig        *YoutubeConfig
}

type AuthenticationConfig struct {
	Host     string
	Port     string
	Protocol string
}

type ProviderConfig struct {
	Port     string
	Protocol string
}

type YoutubeConfig struct {
	*ProviderConfig
	AuthProviderX509CertUrl string
	AuthUri                 string
	ClientId                string
	ClientSecret            string
	ProjectId               string
	RedirectUri             string
	TokenUri                string
}

type SpotifyConfig struct {
	*ProviderConfig
	ClientId     string
	ClientSecret string
	ProjectId    string
}

type DBConfig struct {
	Host     string
	Name     string
	Password string
	Port     string
	SslMode  string
	User     string
}

type FrontConfig struct {
	Host     string
	Port     string
	Protocol string
}

var Cfg *config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func LoadConfig(provider utils.Provider) *config {
	if Cfg != nil {
		fmt.Println("Error loading config: Config already loaded")
		return Cfg
	}

	setDefaultValues(provider)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, using default values:", err)
	}

	Cfg = &config{
		AuthenticationConfig: &AuthenticationConfig{
			Host:     viper.GetString("authentication.host"),
			Port:     viper.GetString("authentication.port"),
			Protocol: viper.GetString("authentication.protocol"),
		},
		DBConfig: &DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     viper.GetString("database.name"),
			SslMode:  viper.GetString("database.sslmode"),
		},
		FrontConfig: &FrontConfig{
			Host:     viper.GetString("front.host"),
			Port:     viper.GetString("front.port"),
			Protocol: viper.GetString("front.protocol"),
		},
		Provider: utils.GetProvider(viper.GetString("provider")),
		SpotifyConfig: &SpotifyConfig{
			ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			ProviderConfig: &ProviderConfig{
				Port:     viper.GetString("spotify.port"),
				Protocol: viper.GetString("spotify.protocol"),
			},
			ProjectId: viper.GetString("spotify.project_id"),
		},
		// TODO: change some of these to be outside of .env
		YoutubeConfig: &YoutubeConfig{
			AuthProviderX509CertUrl: os.Getenv("YOUTUBE_AUTH_PROVIDER_X509_CERT_URL"),
			AuthUri:                 os.Getenv("YOUTUBE_AUTH_URI"),
			ClientId:                os.Getenv("YOUTUBE_CLIENT_ID"),
			ClientSecret:            os.Getenv("YOUTUBE_CLIENT_SECRET"),
			ProviderConfig: &ProviderConfig{
				Port:     viper.GetString("youtube.port"),
				Protocol: viper.GetString("youtube.protocol"),
			},
			ProjectId:   os.Getenv("YOUTUBE_PROJECT_ID"),
			RedirectUri: os.Getenv("YOUTUBE_REDIRECT_URI"),
			TokenUri:    os.Getenv("YOUTUBE_TOKEN_URI"),
		},
	}

	return Cfg
}

func (c *config) GetProvider() *ProviderConfig {
	switch Cfg.Provider {
	case utils.SpotifyProvider:
		return Cfg.SpotifyConfig.ProviderConfig
	case utils.YoutubeProvider:
		return Cfg.YoutubeConfig.ProviderConfig
	default:
		return nil
	}
}

func setDefaultValues(provider utils.Provider) {
	// Provider
	viper.SetDefault("provider", provider.GetString())

	// Spotify port and protocol
	viper.SetDefault("spotify.port", 5001)
	viper.SetDefault("spotify.protocol", "http")

	// Youtube port and protocol
	viper.SetDefault("youtube.port", 5002)
	viper.SetDefault("youtube.protocol", "http")

	// Authentication Config
	viper.SetDefault("Authentication.host", "127.0.0.1")
	viper.SetDefault("Authentication.port", 3333)
	viper.SetDefault("Authentication.protocol", "http")

	// DB Config
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.name", "youtofy")
	viper.SetDefault("database.sslmode", "disable")

	// Front Config
	viper.SetDefault("front.host", "loving-deep-loon.ngrok-free.app")
	viper.SetDefault("front.port", 443)
	viper.SetDefault("front.protocol", "https")
}
