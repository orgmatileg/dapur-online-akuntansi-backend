package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/helper"
	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	PUsecase product.Usecase
}

func NewProductHTTPHandler(r *mux.Router, pu product.Usecase) {

	handler := ProductHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/product", handler.ProductSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/product", handler.ProductFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/product/{id}", handler.ProductFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/product/{id}", handler.ProductUpdateHTTPHandler).Methods("PUT")
	r.HandleFunc("/product/{id}", handler.ProductDeleteHTTPHandler).Methods("DELETE")
	r.HandleFunc("/product/{id}/exists", handler.ProductIsExistsByIDHTTPHandler).Methods("GET")
}

// productTypesSaveHTTPHandler handler
func (u *ProductHandler) ProductSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	m := model.NewProduct()

	err, res := decoder.Decode(m), helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Err, res.Body.Payload = u.PUsecase.Save(m), m

}

// productTypesFindAllHTTPHandler handler
func (u *ProductHandler) ProductFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

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
func (u *ProductHandler) ProductFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	res.Body.Payload, res.Err = u.PUsecase.FindByID(idP)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *ProductHandler) ProductUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var m model.Product

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
func (u *ProductHandler) ProductDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

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
func (u *ProductHandler) ProductIsExistsByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

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
