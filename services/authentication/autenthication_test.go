package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
)

func TestAuthorize(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	configs.LoadConfig()

	// Initialize configs.Cfg and AuthenticationConfig to avoid nil pointer dereference
	if configs.Cfg.AuthenticationConfig == nil {
		configs.Cfg.AuthenticationConfig = &configs.AuthenticationConfig{}
	}
	configs.Cfg.AuthenticationConfig.Protocol = "http"
	configs.Cfg.AuthenticationConfig.Host = mockServer.Listener.Addr().String()
	configs.Cfg.AuthenticationConfig.Port = ""

	tests := []struct {
		name     string
		username string
		cookies  []*http.Cookie
		expected bool
	}{
		{
			name:     "Valid authorization",
			username: "testuser",
			cookies:  []*http.Cookie{{Name: "X-CSRF-Token", Value: "validtoken"}},
			expected: true,
		},
		{
			name:     "Invalid authorization",
			username: "invaliduser",
			cookies:  []*http.Cookie{{Name: "X-CSRF-Token", Value: "invalidtoken"}},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Authorize(tt.username, tt.cookies)
			if result != tt.expected {
				t.Errorf("Authorize() = %v, expected %v", result, tt.expected)
			}
		})
	}
	defer mockServer.Close()
}
