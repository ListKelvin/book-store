package main

import (
	"github.com/ListKelvin/book-store/api"
)




func main(){

	// routes.RegisterBookStoreRoutes(r)
	// http.Handle("/", r)
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// r.Use(mux.CORSMethodMiddleware(r))
	// r.Use(LoggingMiddleware)
	// log.Fatal(http.ListenAndServe("localhost:8080", r))
	api.Run()
}


