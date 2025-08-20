package spotify

type AuthenticationResponse struct {
	Code string `json:"code"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expeires_in"`
	RefreshToken string `json:"refresh_token"`
}
