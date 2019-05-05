package auth

import (
	modelAuth "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/model"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/model"
	"golang.org/x/oauth2"
)

type Usecase interface {
	LoginJWT(*model.User) (*modelAuth.Auth, error)
	Oauth2FacebookLogin() (*oauth2.Config, string)
}
