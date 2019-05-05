package usecase

import (
	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewproductUsecase(r product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: r,
	}
}

func (u *productUsecase) Save(m *model.ProductTypes) (err error) {
	return u.productRepo.Save(m)
}

func (u *productUsecase) FindByID(id string) (m *model.ProductTypes, err error) {
	return u.productRepo.FindByID(id)
}

func (u *productUsecase) FindAll(limit, offset, order string) (ml model.ProductTypesList, count int64, err error) {

	ml, err = u.productRepo.FindAll(limit, offset, order)
	if err != nil {
		return nil, -1, err
	}
	count, err = u.productRepo.Count()

	return
}

func (u *productUsecase) Update(id string, m *model.ProductTypes) (rowAffected *string, err error) {
	return u.productRepo.Update(id, m)
}

func (u *productUsecase) Delete(idUser string) (err error) {
	return u.productRepo.Delete(idUser)
}

func (u *productUsecase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.productRepo.IsExistsByID(idUser)
}

func (u *productUsecase) Count() (count int64, err error) {
	return u.productRepo.Count()
}
