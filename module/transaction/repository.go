package product

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"

// Repository interface
type Repository interface {
	Save(*model.Product) error
	FindByID(id string) (*model.Product, error)
	FindAll(limit, offset, order string) (ml model.ProductList, err error)
	Update(id string, m *model.Product) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
