package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/SOLO-YOLO-BACKEND/helper"
	userRole "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users_role"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users_role/model"

	"github.com/gorilla/mux"
)

type UsersRoleHandler struct {
	UUsecase userRole.Usecase
}

func NewUsersRoleHttpHandler(r *mux.Router, uu userRole.Usecase) {

	handler := UsersRoleHandler{
		UUsecase: uu,
	}

	r.HandleFunc("/users/roles", handler.UserRoleSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/users/roles", handler.UserRoleFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/users/roles/{id}", handler.UserRoleFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/users/roles/{id}", handler.UserRoleUpdateHTTPHandler).Methods("PUT")
	r.HandleFunc("/users/roles/{id}", handler.UserRoleDeleteHTTPHandler).Methods("DELETE")
	r.HandleFunc("/users/roles/{id}/exists", handler.UserRoleIsExistsByIDHTTPHandler).Methods("GET")
}

// UserRoleSaveHTTPHandler handler
func (u *UsersRoleHandler) UserRoleSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	m := model.NewUserRole()

	err, res := decoder.Decode(m), helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Err, res.Body.Payload = u.UUsecase.Save(m), m

}

// UserRoleFindAllHTTPHandler handler
func (u *UsersRoleHandler) UserRoleFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	// Set default query
	limit, offset, order := "10", "0", "desc"

	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}

	if v := queryParam.Get("offset"); v != "" {
		offset = queryParam.Get("offset")
	}

	if v := queryParam.Get("order"); v != "" {
		order = queryParam.Get("order")
	}

	res := helper.Response{}

	res.Body.Payload, res.Body.Count, res.Err = u.UUsecase.FindAll(limit, offset, order)

	res.ServeJSON(w, r)

}

// UserRoleFindByIDHTTPHandler handler
func (u *UsersRoleHandler) UserRoleFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	res.Body.Payload, res.Err = u.UUsecase.FindByID(idP)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *UsersRoleHandler) UserRoleUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var m model.UserRole

	err := decoder.Decode(&m)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.UUsecase.Update(idP, &m)

	if err != nil {
		res.Err = err
		return
	}

	fmt.Println(rowsAffected, err)

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

// UserRoleDeleteHttpHandler handler
func (u *UsersRoleHandler) UserRoleDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP, res := vars["id"], helper.Response{}

	err := u.UUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

// UserRoleIsExistsByIDHttpHandler handler
func (u *UsersRoleHandler) UserRoleIsExistsByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	isExists, err := u.UUsecase.IsExistsByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = struct {
		Exist bool `json:"is_exist"`
	}{
		Exist: isExists,
	}

}
