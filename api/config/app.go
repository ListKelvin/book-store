package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ListKelvin/book-store/api/models"
	"github.com/ListKelvin/book-store/api/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(

	db * gorm.DB
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}


func Connect() {
	dsn := "root:sa123456@tcp(127.0.0.1:3306)/bookStore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&models.User{})
	fmt.Println("Listening to port", "8080")
	db = d
}

func GetDB() *gorm.DB {
	return db
}



var errList = make(map[string]string)

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

		DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
		

	//database migration
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Book{},
		&models.Customer{},
		&models.Genre{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Publisher{},
		&models.Review{},
		&models.Shipper{},
		&models.Discount{},
	)
//Declare route default with mux
	server.Router = mux.NewRouter()
	server.Router.Use(mux.CORSMethodMiddleware(server.Router))

	routes.InitializeRoutes(server.Router)
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}