package model

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"os"
)

type Auth struct {
	UserID       int64  `json:"user_id"`
	NamaLengkap  string `json:"nama_lengkap"`
	PhotoProfile string `json:"photo_profile"`
	Token        string `json:"token"`
}

func NewFacebookOauth2Config() *oauth2.Config {

	facebookAppID := os.Getenv("FB_APP_ID")
	facebookAppSecret := os.Getenv("FB_APP_SECRET")

	return &oauth2.Config{
		ClientID:     facebookAppID,
		ClientSecret: facebookAppSecret,
		RedirectURL:  "http://localhost:8081/v1/oauth/facebook/callback",
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"email"},
	}
}
