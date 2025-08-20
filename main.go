package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	spotiftController "github.com/DevOps-Group-D/YouToFy-API/controllers/spotify"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Reading .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initializing configs
	cfg := configs.LoadConfig()

	// Listening and serving service
	router := chi.NewRouter()

	// Registering get routes
	router.Get("/spotify/login", spotiftController.Login)
	router.Get("/spotify/playlist", spotiftController.GetPlaylist)

	// Registering post routes
	router.Post("/spotify/save", spotiftController.Save)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}
