package database

import (
	"database/sql"
	"fmt"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	cfg := configs.Cfg

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.Name, cfg.DBConfig.SslMode)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	return conn, err
}
