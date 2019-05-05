package router

import (
	"fmt"

	"net/http"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/config"
	// m "github.com/orgmatileg/dapur-online-akuntansi-backend/middleware"

	// Auth
	hAuth "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/delivery/http"
	_authRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/repository"
	_authUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth/usecase"

	// User

	// User Role
	hUserRole "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/delivery/http"
	_usersRoleRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/repository"
	_usersRoleUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/usecase"

	// Product Types
	hProductTypes "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/delivery/http"
	_productTypesRepo "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/repository"
	_productTypesUcase "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/usecase"

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

	// Users Role
	usersRoleRepo := _usersRoleRepo.NewUsersRoleRepositoryPostgres(dbConn)
	usersRoleUcase := _usersRoleUcase.NewUsersRoleUsecase(usersRoleRepo)
	hUserRole.NewUsersRoleHttpHandler(rv1, usersRoleUcase)

	// Product Types
	productTypesRepo := _productTypesRepo.NewProductTypesRepositoryPostgres(dbConn)
	productTypesUcase := _productTypesUcase.NewProductTypesUsecase(productTypesRepo)
	hProductTypes.NewProductTypesHTTPHandler(rv1, productTypesUcase)

	return r
}
