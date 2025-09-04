package spotify

import "github.com/DevOps-Group-D/YouToFy-API/database"

const UPDATE_ACCESS_TOKEN_QUERY = `UPDATE spotify SET access_token = $1 WHERE account_username = $2`

func UpdateAccessToken(username string, accessToken string) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(UPDATE_ACCESS_TOKEN_QUERY, accessToken, username)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
