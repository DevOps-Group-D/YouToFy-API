package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	"github.com/DevOps-Group-D/YouToFy-API/controllers/interfaces"
	"github.com/DevOps-Group-D/YouToFy-API/controllers/spotify"
	"github.com/DevOps-Group-D/YouToFy-API/controllers/youtube"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Reading args from cli
	provider := os.Args[1]

	// Reading .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initializing configs
	cfg := configs.LoadConfig()

	// Listening and serving service
	router := chi.NewRouter()
	router.Use(chimiddleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	providerEnum := utils.GetProvider(provider)
	providerImpl := getProvider(providerEnum)

	// Registering get routes
	router.Get("/login", providerImpl.Login)
	router.Get("/playlist/{playlistId}", providerImpl.GetPlaylist)

	// Registering post routes
	router.Post("/save", providerImpl.Save)

	// Registering patch routes
	router.Patch("/playlist/{playlistId}", providerImpl.InsertPlaylist)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}

func getProvider(provider utils.Provider) interfaces.Provider {
	switch provider {
	case utils.SpotifyProvider:
		return spotify.NewSpotifyProvider()
	case utils.YoutubeProvider:
		return youtube.NewYoutubeProvider()
	default:
		return nil
	}
}
