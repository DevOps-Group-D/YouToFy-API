package authentication

import "database/sql"

type Account struct {
	Username     string         `json:"username"`
	Password     string         `json:"password,omitempty"`
	SessionToken sql.NullString `json:"session_token,omitempty"`
	CsrfToken    sql.NullString `json:"csrf_token,omitempty"`
	Spotify      *Spotify       `json:"spotify"`
}

type Spotify struct {
	AccessToken string `json:"access_token"`
}

type Youtube struct {
	AccountUsername string `json:"account_username"`
	AccessToken     string `json:"access_token"`
}
