package spotify

import (
	spotifyRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/spotify"
	spotifyService "github.com/DevOps-Group-D/YouToFy-API/services/spotify"
)

type spotifyProvider struct {
	Service *spotifyService.SpotifyService
}

func NewSpotifyProvider() *spotifyProvider {
	return &spotifyProvider{
		Service: &spotifyService.SpotifyService{
			Repository: &spotifyRepository.SpotifyRepository{},
		},
	}
}
