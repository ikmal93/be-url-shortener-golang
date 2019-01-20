package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"goland/urlShortenerIkmal/routers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := routers.InitRoutes()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "token"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	fmt.Println("Starting application at http://localhost:12345")
	log.Fatal(http.ListenAndServe("localhost:12345", handlers.CORS(headersOk, methodsOk, originsOk)(router)))
}
