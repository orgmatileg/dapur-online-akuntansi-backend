package auth

import (
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users/model"
)

// Repository interface
type Repository interface {
	LoginJWT(*model.User) (userModel *model.User, err error)
	// Oauth2FacebookLogin() error
}
