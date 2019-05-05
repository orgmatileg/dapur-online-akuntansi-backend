package usecase

import (
	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(r product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: r,
	}
}

func (u *productUsecase) Save(m *model.Product) (err error) {
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
