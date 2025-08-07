func GetWebTokenLink() string {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func loadConfig() (*oauth2.Config, error) {
	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config, nil
}
