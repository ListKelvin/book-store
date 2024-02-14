package controllers

import (
	"fmt"
	"log"

	"github.com/ListKelvin/book-store/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(

	db * gorm.DB
)

type Server struct {
	DB     *gorm.DB

}



func GetDB() *gorm.DB {
	return db
}



var errList = make(map[string]string)

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

		DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		fmt.Print("DbUrl: "+ DbUrl)
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



}

