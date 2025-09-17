package youtube

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
)

func TestGetAuthURL(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/o/oauth2/auth" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Mock Auth URL"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer mockServer.Close()
	configs.LoadConfig()
	configs.Cfg.YoutubeConfig.ClientId = "mock-client-id"
	configs.Cfg.YoutubeConfig.ClientSecret = "mock-client-secret"
	configs.Cfg.YoutubeConfig.RedirectUri = "http://localhost:8080/oauth2callback"
	configs.Cfg.AuthenticationConfig.Protocol = "http"
	configs.Cfg.AuthenticationConfig.Host = "localhost"
	configs.Cfg.AuthenticationConfig.Port = "8080"
	url := GetAuthURL()
	if url == "" {
		t.Errorf("GetAuthURL() returned empty URL")
	}
}

func TestGetPlaylist(t *testing.T) {
	// This test requires a valid OAuth2 token and a real playlist ID.
	// For demonstration, we will skip the actual API call.
	t.Skip("Skipping TestGetPlaylist as it requires real OAuth2 token and playlist ID")
}
