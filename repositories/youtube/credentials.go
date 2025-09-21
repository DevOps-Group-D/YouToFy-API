package youtube

import (
	"github.com/DevOps-Group-D/YouToFy-API/database"
	"github.com/DevOps-Group-D/YouToFy-API/models/authentication"
)

const (
	INSERT_QUERY_YOUTUBECREDENTIALS = `INSERT INTO youtube (account_username, access_token) VALUES ($1, $2)`
	SELECT_QUERY_YOUTUBECREDENTIALS = `SELECT * FROM youtube count WHERE account_username = $1`
	UPDATE_QUERY_YOUTUBECREDENTIALS = `UPDATE youtube SET access_token = $2 WHERE account_username = $1`
)

// TODO: Add a struct to all these methods like Spotify
func InsertYouTubeCredentials(
	Username string,
	AccessToken string,
) error {

	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(
		INSERT_QUERY_YOUTUBECREDENTIALS,
		Username, AccessToken,
	)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func GetYouTubeCredentials(username string) (*authentication.Youtube, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(SELECT_QUERY_YOUTUBECREDENTIALS, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	youTubeCredentials := &authentication.Youtube{}
	err = row.Scan(&youTubeCredentials.AccountUsername, &youTubeCredentials.AccessToken)
	if err != nil {
		return nil, err
	}

	youTubeCredentials.AccountUsername = username

	return youTubeCredentials, nil
}

func UpdateYouTubeCredentials(
	Username string,
	AccessToken string,
) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(UPDATE_QUERY_YOUTUBECREDENTIALS,
		Username, AccessToken)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
