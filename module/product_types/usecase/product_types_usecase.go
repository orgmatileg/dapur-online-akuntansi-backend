package usecase

import (
	productTypes "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/model"
)

type productTypesUsecase struct {
	productTypesRepo productTypes.Repository
}

func NewProductTypesUsecase(r productTypes.Repository) productTypes.Usecase {
	return &productTypesUsecase{
		productTypesRepo: r,
	}
}

func (u *productTypesUsecase) Save(m *model.ProductTypes) (err error) {
	return u.productTypesRepo.Save(m)
}

func (u *productTypesUsecase) FindByID(id string) (m *model.ProductTypes, err error) {
	return u.productTypesRepo.FindByID(id)
}

func (u *productTypesUsecase) FindAll(limit, offset, order string) (ml model.ProductTypesList, count int64, err error) {

	ml, err = u.productTypesRepo.FindAll(limit, offset, order)
	if err != nil {
		return nil, -1, err
	}
	count, err = u.productTypesRepo.Count()

	return
}

func (u *productTypesUsecase) Update(id string, m *model.ProductTypes) (rowAffected *string, err error) {
	return u.productTypesRepo.Update(id, m)
}

func (u *productTypesUsecase) Delete(idUser string) (err error) {
	return u.productTypesRepo.Delete(idUser)
}

func (u *productTypesUsecase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.productTypesRepo.IsExistsByID(idUser)
}

func (u *productTypesUsecase) Count() (count int64, err error) {
	return u.productTypesRepo.Count()
}
