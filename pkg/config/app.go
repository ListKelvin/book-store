package config

import (
	"github.com/ListKelvin/book-store/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



var(

	db * gorm.DB
)

func Connect() {
	dsn := "root:sa123456@tcp(127.0.0.1:3306)/bookStore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	db = d
}

func GetDB() *gorm.DB {
	return db
}