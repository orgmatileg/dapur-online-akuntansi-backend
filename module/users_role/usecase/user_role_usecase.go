package usecase

import (
	userRole "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users_role"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users_role/model"
)

type userRoleUsecase struct {
	userRoleRepo userRole.Repository
}

func NewUsersRoleUsecase(r userRole.Repository) userRole.Usecase {
	return &userRoleUsecase{
		userRoleRepo: r,
	}
}

func (u *userRoleUsecase) Save(m *model.UserRole) (err error) {
	return u.userRoleRepo.Save(m)
}

func (u *userRoleUsecase) FindByID(id string) (m *model.UserRole, err error) {
	return u.userRoleRepo.FindByID(id)
}

func (u *userRoleUsecase) FindAll(limit, offset, order string) (ml model.UserRoleList, count int64, err error) {

	ml, err = u.userRoleRepo.FindAll(limit, offset, order)
	if err != nil {
		return nil, -1, err
	}
	count, err = u.userRoleRepo.Count()

	return
}

func (u *userRoleUsecase) Update(id string, m *model.UserRole) (rowAffected *string, err error) {
	return u.userRoleRepo.Update(id, m)
}

func (u *userRoleUsecase) Delete(idUser string) (err error) {
	return u.userRoleRepo.Delete(idUser)
}

func (u *userRoleUsecase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.userRoleRepo.IsExistsByID(idUser)
}

func (u *userRoleUsecase) Count() (count int64, err error) {
	return u.userRoleRepo.Count()
}
