package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/config"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/router"
)

func init() {

	if env := os.Getenv("GO_ENV"); env != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.InitConnectionDB()
}

func main() {
	router := router.InitRouter()

	originCORS := handlers.AllowedOrigins([]string{"*"})
	headersCORS := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methodCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	fmt.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originCORS, headersCORS, methodCORS)(router)))
}
