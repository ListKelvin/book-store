package models

import (
	"github.com/ListKelvin/book-store/pkg/config"
	// "github.com/ListKelvin/book-store/pkg/models"
	"gorm.io/gorm"
)


var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
    
	result := db.Create(&b)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return b, nil // Return the created book and a nil error if successful
}

func GetAllBooks() []Book{
	Books :=  make([]Book, 0)
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	db:= db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}


func DeleteBook(ID int64) Book {
	var book Book

	 db.Where("ID=?", ID).Delete(book)
	return book
}