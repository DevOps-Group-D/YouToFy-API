package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	ApiConfig            *ApiConfig
	DBConfig             *DBConfig
	FrontConfig          *FrontConfig
	SpotifyConfig        *SpotifyConfig
	AuthenticationConfig *AuthenticationConfig
}

type ApiConfig struct {
	Port     string
	Protocol string
}

type AuthenticationConfig struct {
	Host     string
	Port     string
	Protocol string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
}

type FrontConfig struct {
	Host     string
	Port     string
	Protocol string
}

type SpotifyConfig struct {
	ClientId     string
	ClientSecret string
}

var Cfg *config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func LoadConfig() *config {
	if Cfg != nil {
		fmt.Println("Error loading config: Config already loaded")
		return Cfg
	}

	setDefaultValues()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, using default values:", err)
	}

	Cfg = &config{
		ApiConfig: &ApiConfig{
			Port:     viper.GetString("api.port"),
			Protocol: viper.GetString("api.protocol"),
		},
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
		SpotifyConfig: &SpotifyConfig{
			ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		},
	}

	return Cfg
}

func setDefaultValues() {
	// API Config
	viper.SetDefault("api.port", 3000)
	viper.SetDefault("api.protocol", "http")

	// Authentication Config
	viper.SetDefault("Authentication.host", "127.0.0.1")
	viper.SetDefault("Authentication.port", 3333)
	viper.SetDefault("Authentication.protocol", "http")

	// DB Config
	viper.SetDefault("database.host", "postgres")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.name", "youtofy")
	viper.SetDefault("database.sslmode", "disable")

	// Front Config
	viper.SetDefault("front.host", "127.0.0.1")
	viper.SetDefault("front.port", 8080)
	viper.SetDefault("front.protocol", "http")
}
