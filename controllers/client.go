package controllers

import (
	"net/http"
	"time"
)

var Client *http.Client

func init() {
	Client = &http.Client{
		Timeout: 30 * time.Second,
	}
}
