package youtube

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
)

var mockServer *httptest.Server

func setupTest() {
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/o/oauth2/auth" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Mock Auth URL"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	configs.LoadConfig()
	configs.Cfg.YoutubeConfig.ClientId = "mock-client-id"
	configs.Cfg.YoutubeConfig.ClientSecret = "mock-client-secret"
	configs.Cfg.YoutubeConfig.RedirectUri = "http://localhost:8080/oauth2callback"
	configs.Cfg.AuthenticationConfig.Protocol = "http"
	configs.Cfg.AuthenticationConfig.Host = "localhost"
	configs.Cfg.AuthenticationConfig.Port = "8080"
}

func teardownTest() {
	if mockServer != nil {
		mockServer.Close()
	}
}

func TestGetAuthURL(t *testing.T) {
	setupTest()
	defer teardownTest()
	url := GetAuthURL()
	if url == "" {
		t.Errorf("GetAuthURL() returned empty URL")
	}
}
