package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	ApiConfig *ApiConfig
}

type ApiConfig struct {
	Port string
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
	}

	return Cfg
}

func setDefaultValues() {
	// API Config
	viper.SetDefault("api.port", 3333)
}
