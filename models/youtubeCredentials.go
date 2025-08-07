package models

import "database/sql"

type YouTubeCredentials struct {
	OwnerUsername string         `json:"owner_username"`
	AccessToken   string         `json:"access_token"`
	TokenType     string         `json:"token_type"`
	RefreshToken  sql.NullString `json:"refresh_token"`
	Expiry        sql.NullString `json:"expiry"`
	ExpiresIn     sql.NullInt64  `json:"expires_in"`
}
