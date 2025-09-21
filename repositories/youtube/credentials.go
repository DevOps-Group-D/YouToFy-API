package youtube

import (
	"github.com/DevOps-Group-D/YouToFy-API/database"
)

type YoutubeRepository struct{}

const (
	UPDATE_ACCESS_TOKEN_QUERY = `UPDATE youtube SET access_token = $2 WHERE account_username = $1`
)

func (y *YoutubeRepository) UpdateYouTubeCredentials(
	Username string,
	AccessToken string,
) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(UPDATE_ACCESS_TOKEN_QUERY,
		Username, AccessToken)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
