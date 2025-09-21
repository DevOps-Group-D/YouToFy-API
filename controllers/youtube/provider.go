package youtube

import (
	youtubeRepository "github.com/DevOps-Group-D/YouToFy-API/repositories/youtube"
	youtubeService "github.com/DevOps-Group-D/YouToFy-API/services/youtube"
)

type youtubeProvider struct {
	Service *youtubeService.YoutubeService
}

func NewYoutubeProvider() *youtubeProvider {
	return &youtubeProvider{
		Service: &youtubeService.YoutubeService{
			Repository: &youtubeRepository.YoutubeRepository{},
		},
	}
}
