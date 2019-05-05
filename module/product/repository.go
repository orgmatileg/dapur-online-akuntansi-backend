package product

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"

// Repository interface
type Repository interface {
	Save(*model.ProductTypes) error
	FindByID(id string) (*model.ProductTypes, error)
	FindAll(limit, offset, order string) (mel model.ProductTypesList, err error)
	Update(id string, m *model.ProductTypes) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
