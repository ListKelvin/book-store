package routes

import (
	"net/http"

	"github.com/ListKelvin/book-store/api/controllers"
	"github.com/ListKelvin/book-store/api/utils"

	"github.com/gorilla/mux"
)



var RegisterBookStoreRoutes = func(router *mux.Router) { 
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost,  http.MethodOptions)
	router.HandleFunc("/book/", controllers.GetBooks).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(http.MethodGet,  http.MethodOptions)
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods( http.MethodPut, http.MethodOptions)
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods( http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/auth/google/login", utils.OauthGoogleLogin)
	router.HandleFunc("/auth/google/callback", utils.OauthGoogleCallback)
}