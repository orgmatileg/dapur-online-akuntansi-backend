package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/helper"
	productTypes "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/model"

	"github.com/gorilla/mux"
)

type ProductTypesHandler struct {
	PUsecase productTypes.Usecase
}

func NewProductTypesHTTPHandler(r *mux.Router, pu productTypes.Usecase) {

	handler := ProductTypesHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/product-types", handler.ProductTypesSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/product-types", handler.ProductTypesFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/product-types/{id}", handler.ProductTypesFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/product-types/{id}", handler.ProductTypesUpdateHTTPHandler).Methods("PUT")
	r.HandleFunc("/product-types/{id}", handler.ProductTypesDeleteHTTPHandler).Methods("DELETE")
	r.HandleFunc("/product-types/{id}/exists", handler.ProductTypesIsExistsByIDHTTPHandler).Methods("GET")
}

// productTypesSaveHTTPHandler handler
func (u *ProductTypesHandler) ProductTypesSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	m := model.NewProductTypes()

	err, res := decoder.Decode(m), helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Err, res.Body.Payload = u.PUsecase.Save(m), m

}

// productTypesFindAllHTTPHandler handler
func (u *ProductTypesHandler) ProductTypesFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

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

	res.Body.Payload, res.Body.Count, res.Err = u.PUsecase.FindAll(limit, offset, order)

	res.ServeJSON(w, r)

}

// productTypesFindByIDHTTPHandler handler
func (u *ProductTypesHandler) ProductTypesFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	res.Body.Payload, res.Err = u.PUsecase.FindByID(idP)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *ProductTypesHandler) ProductTypesUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var m model.ProductTypes

	err := decoder.Decode(&m)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.PUsecase.Update(idP, &m)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

// productTypesDeleteHttpHandler handler
func (u *ProductTypesHandler) ProductTypesDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP, res := vars["id"], helper.Response{}

	err := u.PUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

// productTypesIsExistsByIDHttpHandler handler
func (u *ProductTypesHandler) ProductTypesIsExistsByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	isExists, err := u.PUsecase.IsExistsByID(idP)

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
