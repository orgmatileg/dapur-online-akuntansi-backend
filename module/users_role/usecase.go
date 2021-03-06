package users_role

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/model"

type Usecase interface {
	Save(*model.UserRole) error
	FindByID(id string) (*model.UserRole, error)
	FindAll(limit, offset, order string) (ml model.UserRoleList, count int64, err error)
	Update(id string, m *model.UserRole) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
