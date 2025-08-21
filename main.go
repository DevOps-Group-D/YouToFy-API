package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	spotiftController "github.com/DevOps-Group-D/YouToFy-API/controllers/spotify"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	router.Use(chimiddleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Registering get routes
	router.Get("/spotify/login", spotiftController.Login)
	router.Get("/spotify/playlist/{playlistId}", spotiftController.GetPlaylist)

	// Registering post routes
	router.Post("/spotify/save", spotiftController.Save)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}
