package example

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/_example/model"

// Repository interface
type Repository interface {
	Save(*model.Example) error
	FindByID(id string) (*model.Example, error)
	FindAll(limit, offset, order string) (mel model.ExampleList, err error)
	Update(id string, modelUser *model.Example) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
