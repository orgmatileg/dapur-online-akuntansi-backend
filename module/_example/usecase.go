package example

import "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/_example/model"

type Usecase interface {
	Save(*model.Example) error
	FindByID(id string) (*model.Example, error)
	FindAll(limit, offset, order string) (mel model.ExampleList, count int64, err error)
	Update(id string, exampleModel *model.Example) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
