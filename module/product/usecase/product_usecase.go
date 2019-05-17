package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"
	"google.golang.org/api/option"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(r product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: r,
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

func (u *productUsecase) Save(m *model.Product) (err error) {
	if m.Image != "" {
		config := &firebase.Config{
			StorageBucket: "dapur-online.appspot.com",
		}

		secretFile := "./dapur-online-firebase-adminsdk-2m3s6-dfdae8ffb2.json"
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

		sw := bucket.Object(fmt.Sprintf("produk-penjualan/produk-%d.png", time.Now().Unix())).NewWriter(context.Background())
		sw.ContentType = "image/png"

		b, err := base64towriter(m.Image)
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

		googleAPIURL := "https://www.googleapis.com/download/storage/v1/b/dapur-online.appspot.com/o/produk-penjualan"
		firebaseAPIURL := "https://firebasestorage.googleapis.com/v0/b/dapur-online.appspot.com/o/produk-penjualan"
		m.Image = strings.Replace(sw.Attrs().MediaLink, googleAPIURL, firebaseAPIURL, -1)

		fmt.Println(i, err)
	}

	return u.productRepo.Save(m)
}

func (u *productUsecase) FindByID(id string) (m *model.Product, err error) {
	return u.productRepo.FindByID(id)
}

func (u *productUsecase) FindAll(limit, offset, order string) (ml model.ProductList, count int64, err error) {
	ml, err = u.productRepo.FindAll(limit, offset, order)
	if err != nil {
		return nil, -1, err
	}
	count, err = u.productRepo.Count()
	return
}

func (u *productUsecase) Update(id string, m *model.Product) (rowAffected *string, err error) {
	if m.Image == "" {
		v, err := u.FindByID(id)
		if err != nil {
			return nil, err
		}
		m.Image = v.Image
	} else {
		config := &firebase.Config{
			StorageBucket: "dapur-online.appspot.com",
		}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		secretFile := dir + "/module/product/usecase/dapur-online-firebase-adminsdk-2m3s6-dfdae8ffb2.json"
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

		sw := bucket.Object(fmt.Sprintf("produk-penjualan/produk-%d.png", time.Now().Unix())).NewWriter(context.Background())
		sw.ContentType = "image/png"

		b, err := base64towriter(m.Image)
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

		googleAPIURL := "https://www.googleapis.com/download/storage/v1/b/dapur-online.appspot.com/o/produk-penjualan"
		firebaseAPIURL := "https://firebasestorage.googleapis.com/v0/b/dapur-online.appspot.com/o/produk-penjualan"
		m.Image = strings.Replace(sw.Attrs().MediaLink, googleAPIURL, firebaseAPIURL, -1)

		fmt.Println(i, err)

	}
	return u.productRepo.Update(id, m)
}

func (u *productUsecase) Delete(id string) (err error) {
	return u.productRepo.Delete(id)
}

func (u *productUsecase) IsExistsByID(id string) (isExist bool, err error) {
	return u.productRepo.IsExistsByID(id)
}

func (u *productUsecase) Count() (count int64, err error) {
	return u.productRepo.Count()
}

// func (u *productUsecase) Count() (count int64, err error) {
// 	return u.productRepo.Count()
// }
