package repositoriesAcc

import (
	"github.com/DevOps-Group-D/YouToFy-API/database"
	"github.com/DevOps-Group-D/YouToFy-API/models"
)

const (
	INSERT_QUERY_YOUTUBECREDENTIALS = `INSERT INTO youtube_credentials (owner_username, access_token, token_type, refresh_token, expiry, expires_in) VALUES ($1, $2, $3, $4, $5, $6)`
	SELECT_QUERY_YOUTUBECREDENTIALS = `SELECT * FROM acyoutube_credentials count WHERE username = $1`
	UPDATE_QUERY_YOUTUBECREDENTIALS = `UPDATE youtube_credentials SET access_token = $2, token_type = $3, refresh_token = $4, expiry = $5, expires_in = $6 WHERE owner_username = $1`
)

func InsertYouTubeCredentials(
	Username string,
	AccessToken string,
	TokenType string,
	RefreshToken string,
	Expiry string,
	ExpiresIn string) error {

	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(
		INSERT_QUERY_YOUTUBECREDENTIALS,
		Username, AccessToken,
		TokenType, RefreshToken,
		Expiry, ExpiresIn)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func GetYouTubeCredentials(username string) (*models.YouTubeCredentials, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(SELECT_QUERY_YOUTUBECREDENTIALS, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	youTubeCredentials := &models.YouTubeCredentials{}
	err = row.Scan(
		&youTubeCredentials.AccessToken, &youTubeCredentials.TokenType,
		&youTubeCredentials.RefreshToken, &youTubeCredentials.Expiry,
		&youTubeCredentials.ExpiresIn)
	if err != nil {
		return nil, err
	}

	return youTubeCredentials, nil
}

func UpdateYouTubeCredentials(youTubeCredentials *models.YouTubeCredentials) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(UPDATE_QUERY_YOUTUBECREDENTIALS,
		youTubeCredentials.OwnerUsername, youTubeCredentials.AccessToken,
		youTubeCredentials.TokenType, youTubeCredentials.RefreshToken.String,
		youTubeCredentials.Expiry.String, youTubeCredentials.ExpiresIn.Int64)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
