package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/model"
	"google.golang.org/api/option"

	"time"

	"golang.org/x/crypto/bcrypt"
)

type usersUsercase struct {
	usersRepo users.Repository
}

func NewUsersUsecase(u users.Repository) users.Usecase {
	return &usersUsercase{
		usersRepo: u,
	}
}

func base64towriter(b64 string) ([]byte, error) {

	indexNum := strings.Index(string(b64), ",")
	stringSplit := string(b64)[indexNum+1:]
	unbased, err := base64.StdEncoding.DecodeString(string(stringSplit))

	if err != nil {
		return nil, err
	}

	return unbased, nil

}

func (u *usersUsercase) Save(mu *model.User) (err error) {

	// Handle Photo Profile
	defaultPhotoProfile := "https://i.ibb.co/whHn3mf/default-photo-profile.png"

	if mu.PhotoProfile != "" {
		config := &firebase.Config{
			StorageBucket: "dapur-online.appspot.com",
		}

		secretFile := "/firebase-token.json"
		fmt.Println(secretFile)
		opt := option.WithCredentialsFile(secretFile)

		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalln(err, "1")
		}
		client, err := app.Storage(context.Background())
		if err != nil {
			log.Fatalln(err, "2")
		}

		bucket, err := client.DefaultBucket()
		if err != nil {
			log.Fatalln(err, "3")
		}

		sw := bucket.Object(fmt.Sprintf("user-photo-profile/photo-profile-%d.png", time.Now().Unix())).NewWriter(context.Background())
		sw.ContentType = "image/png"

		b, err := base64towriter(mu.PhotoProfile)
		if err != nil {
			fmt.Println(err)
		}

		i, err := sw.Write(b)

		if err != nil {
			fmt.Println(err, "5")
		}

		err = sw.Close()

		if err != nil {
			fmt.Println(err, "6")
		}

		fmt.Println(sw.MediaLink, "url gambar")

		googleAPIURL := "https://www.googleapis.com/download/storage/v1/b/dapur-online.appspot.com/o/user-photo-profile"
		firebaseAPIURL := "https://firebasestorage.googleapis.com/v0/b/dapur-online.appspot.com/o/user-photo-profile"
		mu.PhotoProfile = strings.Replace(sw.Attrs().MediaLink, googleAPIURL, firebaseAPIURL, -1)

		fmt.Println(i, err)
	} else {
		mu.PhotoProfile = defaultPhotoProfile
	}

	// Handle Password - hashmac
	hash, err := bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	mu.Password = string(hash)

	err = u.usersRepo.Save(mu)

	return err
}

func (u *usersUsercase) FindByID(idUser string) (mu *model.User, err error) {

	mu, err = u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	// Clean password, so password not exposed to public
	mu.Password = ""

	return mu, nil
}

func (u *usersUsercase) FindAll(limit, offset, order string) (mul model.Users, count int64, err error) {

	mul, err = u.usersRepo.FindAll(limit, offset, order)

	if err != nil {
		return nil, -1, err
	}

	count, err = u.usersRepo.Count()

	if err != nil {
		return nil, -1, err
	}

	// Clean password, so password not exposed to public
	for _, v := range mul {
		v.Password = ""
	}

	return
}

func (u *usersUsercase) Update(idUser string, mu *model.User) (rowAffected *string, err error) {

	v, err := u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	if mu.PhotoProfile != "" {
		config := &firebase.Config{
			StorageBucket: "dapur-online.appspot.com",
		}

		secretFile := "/firebase-token.json"
		fmt.Println(secretFile)
		opt := option.WithCredentialsFile(secretFile)

		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalln(err, "1")
		}
		client, err := app.Storage(context.Background())
		if err != nil {
			log.Fatalln(err, "2")
		}

		bucket, err := client.DefaultBucket()
		if err != nil {
			log.Fatalln(err, "3")
		}

		sw := bucket.Object(fmt.Sprintf("user-photo-profile/photo-profile-%d.png", time.Now().Unix())).NewWriter(context.Background())
		sw.ContentType = "image/png"

		b, err := base64towriter(mu.PhotoProfile)
		if err != nil {
			fmt.Println(err)
		}

		i, err := sw.Write(b)

		if err != nil {
			fmt.Println(err, "5")
		}

		err = sw.Close()

		if err != nil {
			fmt.Println(err, "6")
		}

		fmt.Println(sw.MediaLink, "url gambar")

		googleAPIURL := "https://www.googleapis.com/download/storage/v1/b/dapur-online.appspot.com/o/user-photo-profile"
		firebaseAPIURL := "https://firebasestorage.googleapis.com/v0/b/dapur-online.appspot.com/o/user-photo-profile"
		mu.PhotoProfile = strings.Replace(sw.Attrs().MediaLink, googleAPIURL, firebaseAPIURL, -1)

		fmt.Println(i, err)
	} else {
		mu.PhotoProfile = v.PhotoProfile
	}

	// Handle Password
	if mu.Password != "" {
		// create hashmac for the password
		hash, err := bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		mu.Password = string(hash)
	} else {
		// if password is empty, use password user from db
		mu.Password = v.Password
	}

	mu.UpdatedAt = time.Now()

	rowAffected, err = u.usersRepo.Update(idUser, mu)

	if err != nil {
		return nil, err
	}

	return rowAffected, err
}

func (u *usersUsercase) Delete(idUser string) (err error) {

	err = u.usersRepo.Delete(idUser)

	return err
}

func (u *usersUsercase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.usersRepo.IsExistsByID(idUser)
}

func (u *usersUsercase) Count() (count int64, err error) {
	return u.usersRepo.Count()
}
