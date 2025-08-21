package tests

import (
	// error package in erros/account.go
	"errors"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/database"
	yterrors "github.com/DevOps-Group-D/YouToFy-API/errors"
	"github.com/DevOps-Group-D/YouToFy-API/models"
	ServiceAcc "github.com/DevOps-Group-D/YouToFy-API/services"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

func TestRegister(t *testing.T) {
	accountData := struct {
		username string
		Password string
	}{
		username: "testuser",
		Password: "testpassword",
	}
	err := ServiceAcc.Register(accountData.username, accountData.Password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}

func TestGetUser(t *testing.T) {
	accountData := struct {
		username string
		Password string
	}{
		username: "testuser",
		Password: "testpassword",
	}
	err := ServiceAcc.Register(accountData.username, accountData.Password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}
	conn, err := database.Connect()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()
	row := conn.QueryRow("SELECT * FROM account WHERE username = $1", accountData.username)
	if row.Err() != nil {
		t.Fatalf("failed to query user: %v", row.Err())
	}
	account := &models.Account{}
	err = row.Scan(&account.Username, &account.Password, &account.SessionToken, &account.CsrfToken)
	if err != nil {
		t.Fatalf("failed to scan user: %v", err)
	}
	hashedPassword, err := utils.HashPassword(accountData.username)
	if err != nil {
		t.Errorf("error hashing password")
	}
	if account.Username != hashedPassword {
		t.Errorf("expected username %s, got %s", accountData.username, account.Username)
	}

}
