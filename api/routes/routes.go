package routes

import (
	"github.com/ListKelvin/book-store/api/controllers"
	"github.com/gorilla/mux"
)

// var RegisterBookStoreRoutes = func(router *mux.Router) {
// 	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost,  http.MethodOptions)
// 	router.HandleFunc("/book/", controllers.GetBooks).Methods(http.MethodGet, http.MethodOptions)
// 	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(http.MethodGet,  http.MethodOptions)
// 	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods( http.MethodPut, http.MethodOptions)
// 	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods( http.MethodDelete, http.MethodOptions)
// 	router.HandleFunc("/auth/google/login", utils.OauthGoogleLogin)
// 	router.HandleFunc("/auth/google/callback", utils.OauthGoogleCallback)
// }



func InitializeRoutes(s *mux.Router) {
	// s.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	v1 := s.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/book/", controllers.GetBooks)
	// v1.HandleFunc("/book/", controllers.CreateBook)
	
}