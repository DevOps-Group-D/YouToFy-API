package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	ApiConfig     *ApiConfig
	DBConfig      *DBConfig
	YoutubeConfig *YoutubeConfig
}

type ApiConfig struct {
	Port string
}

type YoutubeConfig struct {
	ClientId                string
	ProjectId               string
	AuthUri                 string
	TokenUri                string
	AuthProviderX509CertUrl string
	ClientSecret            string
	RedirectUri             string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
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
			Port: viper.GetString("api.port"),
		},
		DBConfig: &DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     viper.GetString("database.name"),
			SslMode:  viper.GetString("database.sslmode"),
		},
		YoutubeConfig: &YoutubeConfig{
			ClientId:                os.Getenv("YOUTUBE_CLIENT_ID"),
			ProjectId:               os.Getenv("YOUTUBE_PROJECT_ID"),
			AuthUri:                 os.Getenv("YOUTUBE_AUTH_URI"),
			TokenUri:                os.Getenv("YOUTUBE_TOKEN_URI"),
			AuthProviderX509CertUrl: os.Getenv("YOUTUBE_AUTH_PROVIDER_X509_CERT_URL"),
			ClientSecret:            os.Getenv("YOUTUBE_CLIENT_SECRET"),
			RedirectUri:             os.Getenv("YOUTUBE_REDIRECT_URI"),
		},
	}

	return Cfg
}

func setDefaultValues() {
	// API Config
	viper.SetDefault("api.port", 3333)

	// DB Config
	viper.SetDefault("database.host", "postgres")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.name", "youtofy")
	viper.SetDefault("database.sslmode", "disable")
}
