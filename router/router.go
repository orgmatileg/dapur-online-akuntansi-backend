package router

import (
	"fmt"

	"net/http"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/config"

	// Auth
	hAuth "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/delivery/http"
	_authRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/repository"
	_authUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/usecase"

	// User
	hUser "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/delivery/http"
	_usersRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/repository"
	_usersUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/usecase"

	// User Role
	hUserRole "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/delivery/http"
	_usersRoleRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/repository"
	_usersRoleUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/usecase"

	// Product Types
	hProductTypes "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/delivery/http"
	_productTypesRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/repository"
	_productTypesUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/usecase"

	// Product
	hProduct "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/delivery/http"
	_productRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/repository"
	_productUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/usecase"

	// Transaction
	hTransaction "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/delivery/http"
	_transactionRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/repository"
	_transactionUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/usecase"

	"github.com/gorilla/mux"
)

// InitRouter endpoint
func InitRouter() *mux.Router {

	r := mux.NewRouter()
	// Check API
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}).Methods("GET")
	// Endpoint for testing app or such a thing
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Test!")
	}).Methods("POST")

	// Init versioning API
	rv1 := r.PathPrefix("/v1").Subrouter()

	// Middleware
	// rv1.Use(m.CheckAuth)

	// Get DB Conn
	dbConn := config.GetPostgresDB()

	// Auth
	authRepo := _authRepo.NewAuthRepositoryPostgres(dbConn)
	authUcase := _authUcase.NewAuthUsecase(authRepo)
	hAuth.NewAuthHttpHandler(rv1, authUcase)

	// Users
	usersRepo := _usersRepo.NewUserRepositoryPostgres(dbConn)
	usersUcase := _usersUcase.NewUsersUsecase(usersRepo)
	hUser.NewUsersHttpHandler(rv1, usersUcase)

	// Users Role
	usersRoleRepo := _usersRoleRepo.NewUsersRoleRepositoryPostgres(dbConn)
	usersRoleUcase := _usersRoleUcase.NewUsersRoleUsecase(usersRoleRepo)
	hUserRole.NewUsersRoleHttpHandler(rv1, usersRoleUcase)

	// Product Types
	productTypesRepo := _productTypesRepo.NewProductTypesRepositoryPostgres(dbConn)
	productTypesUcase := _productTypesUcase.NewProductTypesUsecase(productTypesRepo)
	hProductTypes.NewProductTypesHTTPHandler(rv1, productTypesUcase)

	// Product
	productRepo := _productRepo.NewProductRepositoryPostgres(dbConn)
	productUcase := _productUcase.NewProductUsecase(productRepo)
	hProduct.NewProductHTTPHandler(rv1, productUcase)

	// Transaction
	transactionRepo := _transactionRepo.NewTransactionRepositoryPostgres(dbConn)
	transactionUcase := _transactionUcase.NewTransactionUsecase(transactionRepo)
	hTransaction.NewTransactionHTTPHandler(rv1, transactionUcase)

	return r
}
