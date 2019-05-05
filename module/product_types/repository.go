package product_types

import "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/product_types/model"

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
