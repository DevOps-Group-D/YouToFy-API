package spotify

type AuthenticationRequest struct {
	Code  string `json:"code"`
	Error string `json:"error"`
	State string `json:"state"`
}
