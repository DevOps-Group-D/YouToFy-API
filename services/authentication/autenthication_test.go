package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

var mockServer *httptest.Server

func setupTest() {
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/authorize" {
			if r.Header.Get("X-CSRF-Token") == "validtoken" {
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}))

	configs.LoadConfig(utils.GetProvider("youtube"))

	// Initialize configs.Cfg and AuthenticationConfig to avoid nil pointer dereference
	if configs.Cfg.AuthenticationConfig == nil {
		configs.Cfg.AuthenticationConfig = &configs.AuthenticationConfig{}
	}
	configs.Cfg.AuthenticationConfig.Protocol = "http"
	configs.Cfg.AuthenticationConfig.Host = mockServer.Listener.Addr().String()
	configs.Cfg.AuthenticationConfig.Port = ""
}

func teardownTest() {
	if mockServer != nil {
		mockServer.Close()
	}
}

func TestAuthorize(t *testing.T) {
	setupTest()
	defer teardownTest()
	result := Authorize("testuser", []*http.Cookie{{Name: "X-CSRF-Token", Value: "validtoken"}})
	if result != true {
		t.Errorf("Authorize() = %v, expected true", result)
	}
}

func TestAuthorizeInvalid(t *testing.T) {
	setupTest()
	defer teardownTest()
	result := Authorize("invaliduser", []*http.Cookie{{Name: "X-CSRF-Token", Value: "invalidtoken"}})
	if result != false {
		t.Errorf("Authorize() = %v, expected false", result)
	}
}
