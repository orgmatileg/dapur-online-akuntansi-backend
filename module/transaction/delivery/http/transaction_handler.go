package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/helper"
	transaction "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/model"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	TUsecase transaction.Usecase
}

func NewTransactionHTTPHandler(r *mux.Router, tu transaction.Usecase) {

	handler := TransactionHandler{
		TUsecase: tu,
	}

	r.HandleFunc("/transaction", handler.TransactionSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/transaction", handler.TransactionFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/transaction/{id}", handler.TransactionFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/transaction/{id}", handler.TransactionUpdateHTTPHandler).Methods("PUT")
	r.HandleFunc("/transaction/{id}", handler.TransactionDeleteHTTPHandler).Methods("DELETE")
	r.HandleFunc("/transaction/{id}/exists", handler.TransactionIsExistsByIDHTTPHandler).Methods("GET")
}

// productTypesSaveHTTPHandler handler
func (u *TransactionHandler) TransactionSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	m := model.NewTransaction()

	err, res := decoder.Decode(m), helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Err, res.Body.Payload = u.TUsecase.Save(m), m

}

// productTypesFindAllHTTPHandler handler
func (u *TransactionHandler) TransactionFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

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

	res.Body.Payload, res.Body.Count, res.Err = u.TUsecase.FindAll(limit, offset, order)

	res.ServeJSON(w, r)

}

// productTypesFindByIDHTTPHandler handler
func (u *TransactionHandler) TransactionFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	res.Body.Payload, res.Err = u.TUsecase.FindByID(idP)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *TransactionHandler) TransactionUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var m model.Transaction

	err := decoder.Decode(&m)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.TUsecase.Update(idP, &m)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

// productTypesDeleteHttpHandler handler
func (u *TransactionHandler) TransactionDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP, res := vars["id"], helper.Response{}

	err := u.TUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

// productTypesIsExistsByIDHttpHandler handler
func (u *TransactionHandler) TransactionIsExistsByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, idP := helper.Response{}, vars["id"]

	isExists, err := u.TUsecase.IsExistsByID(idP)

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
