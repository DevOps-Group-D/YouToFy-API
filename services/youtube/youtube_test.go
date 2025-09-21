package youtube

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	"github.com/DevOps-Group-D/YouToFy-API/repositories/youtube"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

var mockServer *httptest.Server
var youtubeService *YoutubeService

func setupTest() {
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/o/oauth2/auth" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Mock Auth URL"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	youtubeService = &YoutubeService{&youtube.YoutubeRepository{}}
	configs.LoadConfig(utils.GetProvider("youtube"))
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
	url := youtubeService.GetAuthURL()
	if url == "" {
		t.Errorf("GetAuthURL() returned empty URL")
	}
}
