package main

import (
	"log"
	"net/http"

	"github.com/ListKelvin/book-store/pkg/config"
	"github.com/ListKelvin/book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func init() {
	config.Connect()

	server = gin.Default()
}
func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	
}


