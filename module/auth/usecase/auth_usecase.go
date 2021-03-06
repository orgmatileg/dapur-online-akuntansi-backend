package usecase

import (
	"errors"
	"time"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/helper"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth"
	modelAuth "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/model"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/model"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepo auth.Repository
}

func NewAuthUsecase(a auth.Repository) auth.Usecase {
	return &authUsecase{
		authRepo: a,
	}
}

func (a *authUsecase) LoginJWT(mu *model.User) (ma *modelAuth.Auth, err error) {

	userFromDB, err := a.authRepo.LoginJWT(mu)

	if err != nil {
		return nil, err
	}

	// Jika user berhasil ditemukan, maka langkah selanjutnya adalah memvalidasi password
	// hash yang ada di dalam database dengan inputan yang dikirimkan
	if e := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(mu.Password)); e != nil {
		err = errors.New("Password yang Anda masukkan salah!")
		return
	}

	claims := helper.JWTPayload{
		StandardClaims: &jwt.StandardClaims{
			Audience:  "General",
			Issuer:    "Luqmanul Hakim API",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(1440)).Unix(),
		},
	}

	jwtGen := helper.GetJWTTokenGenerator()
	token, err := jwtGen.GenerateToken(claims)

	payload := modelAuth.Auth{
		UserID:       userFromDB.UserID,
		NamaLengkap:  userFromDB.FirstName + " " + userFromDB.LastName,
		PhotoProfile: userFromDB.PhotoProfile,
		Token:        token,
	}

	ma = &payload

	return ma, err
}

// func (a *authUsecase) Oauth2FacebookLogin() (*oauth2.Config, string) {
// 	return modelAuth.NewFacebookOauth2Config(), uuid.NewV4().String()
// }
