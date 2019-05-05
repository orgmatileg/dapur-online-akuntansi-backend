package users_role

import "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users_role/model"

// Repository interface
type Repository interface {
	Save(*model.UserRole) error
	FindByID(id string) (*model.UserRole, error)
	FindAll(limit, offset, order string) (mel model.UserRoleList, err error)
	Update(id string, m *model.UserRole) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
