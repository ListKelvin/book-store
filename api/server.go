package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ListKelvin/book-store/api/controllers"
	"github.com/ListKelvin/book-store/api/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
// init is a function that is called automatically when the program starts
// it loads the environment variables from the .env file

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// This is for testing, when done, do well to comment
	// seed.Load(server.DB)
// old school
	// apiPort := fmt.Sprintf("127.0.0.1:%s ", os.Getenv("API_PORT"))

	fmt.Printf("Listening to port %s", os.Getenv("API_PORT"))


	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	log.Fatal(http.ListenAndServe(os.Getenv("API_PORT"), router))
}