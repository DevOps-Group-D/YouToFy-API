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
