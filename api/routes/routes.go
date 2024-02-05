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


var InitializeRoutes = func(S *mux.Router) {

	v1 := S.PathPrefix("/api/v1").Subrouter()
	{
		// Book Route
	v1.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost,  http.MethodOptions)


		// // Reset password:
		// v1.HandleFunc("/password/forgot", s.ForgotPassword)
		// v1.HandleFunc("/password/reset", s.ResetPassword)

		// //Users routes
		// v1.HandleFunc("/users", s.CreateUser)
		// The user of the app have no business getting all the users.
		// v1.GET("/users", s.GetUsers)
		// v1.GET("/users/:id", s.GetUser)
		// v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		// v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		// v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		// //Posts routes
		// v1.POST("/posts", middlewares.TokenAuthMiddleware(), s.CreatePost)
		// v1.GET("/posts", s.GetPosts)
		// v1.GET("/posts/:id", s.GetPost)
		// v1.PUT("/posts/:id", middlewares.TokenAuthMiddleware(), s.UpdatePost)
		// v1.DELETE("/posts/:id", middlewares.TokenAuthMiddleware(), s.DeletePost)
		// v1.GET("/user_posts/:id", s.GetUserPosts)

		// //Like route
		// v1.GET("/likes/:id", s.GetLikes)
		// v1.POST("/likes/:id", middlewares.TokenAuthMiddleware(), s.LikePost)
		// v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), s.UnLikePost)

		// //Comment routes
		// v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), s.CreateComment)
		// v1.GET("/comments/:id", s.GetComments)
		// v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), s.UpdateComment)
		// v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(), s.DeleteComment)
	}
}