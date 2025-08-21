package servicesAcc

import (
	// error package in erros/account.go
	"errors"
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/database"
	yterrors "github.com/DevOps-Group-D/YouToFy-API/errors"
	"github.com/DevOps-Group-D/YouToFy-API/models"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

func TestLogin(t *testing.T) {
	accountData := struct {
		username string
		Password string
	}{
		username: "testuser",
		Password: "testpassword",
	}
	err := Register(accountData.username, accountData.Password)
	if err != nil {
		t.Fatalf("failed to register user: %v", err)
	}

	account, err := Login(accountData.username, accountData.Password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if account.Username != accountData.username {
		t.Errorf("expected username %s, got %s", accountData.username, account.Username)
	}
}

func TestRegister(t *testing.T) {
	accountData := struct {
		username string
		Password string
	}{
		username: "testuser",
		Password: "testpassword",
	}
	err := Register(accountData.username, accountData.Password)
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
	err := Register(accountData.username, accountData.Password)
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

func TestUnauthorizedError_Error(t *testing.T) {
	err := &yterrors.UnauthorizedError{}
	expected := "UnauthorizedError"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestUnauthorizedError_IsError(t *testing.T) {
	var err error = &yterrors.UnauthorizedError{}
	if !errors.Is(err, &yterrors.UnauthorizedError{}) {
		t.Error("UnauthorizedError should be recognized as itself using errors.Is")
	}
}
