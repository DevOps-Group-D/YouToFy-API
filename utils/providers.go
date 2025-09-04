package utils

type Provider int

const (
	IdleProvider Provider = iota
	SpotifyProvider
	YoutubeProvider
)

var providerNames = []string{"idle", "spotify", "youtube"}
var providers = map[string]Provider{"idle": IdleProvider, "spotify": SpotifyProvider, "youtube": YoutubeProvider}

func (p Provider) GetString() string {
	return providerNames[p]
}

func GetProvider(s string) Provider {
	return providers[s]
}
