package tests

import (
	"testing"

	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

func TestHash(t *testing.T) {
	password := "testpassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if hashedPassword == "" {
		t.Error("expected hashed password to be non-empty")
	}
	err = utils.CheckHashedPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("expected no error when checking hashed password, got %v", err)
	}
}
