package product_types

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/model"

type Usecase interface {
	Save(*model.ProductTypes) error
	FindByID(id string) (*model.ProductTypes, error)
	FindAll(limit, offset, order string) (mel model.ProductTypesList, count int64, err error)
	Update(id string, m *model.ProductTypes) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
