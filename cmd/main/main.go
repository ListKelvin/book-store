package main

import (
	"log"
	"net/http"

	_ "github.com/ListKelvin/book-store/pkg/middleware"
	"github.com/ListKelvin/book-store/pkg/routes"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

// func init() {
// 	config.Connect()

// }
func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// r.Use(mux.CORSMethodMiddleware(r))
	// r.Use(LoggingMiddleware)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	
}


