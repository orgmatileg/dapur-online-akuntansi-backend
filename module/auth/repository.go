package auth

import (
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/model"
)

// Repository interface
type Repository interface {
	LoginJWT(*model.User) (userModel *model.User, err error)
	// Oauth2FacebookLogin() error
}
