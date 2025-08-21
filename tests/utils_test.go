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

func TestToken(t *testing.T) {
	token, err := utils.GenerateToken(32)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(token) != 32 {
		t.Errorf("expected token length to be 32, got %d", len(token))
	}
	if token == "" {
		t.Error("expected token to be non-empty")
	}
	otherToken, err := utils.GenerateToken(32)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token == otherToken {
		t.Error("expected tokens to be unique")
	}
}
