package models

import (
	"time"

	"github.com/ListKelvin/book-store/pkg/config"
	// "github.com/ListKelvin/book-store/pkg/models"
	"gorm.io/gorm"
)


var db *gorm.DB

type Book struct {
	gorm.Model
	Isbn string `gorm:"primary_key;size:255;not null;unique" json:"isbn"`
	Title string `gorm:"size:255;not null;unique" json:"title,omitempty"`
	Available_Quantity uint16 `gorm:"default:0" json:"available_quantity"`
	Price uint16 `gorm:"default:0" json:"price"`
	Edition uint8 `gorm:"default:0" json:"edition"`
	Publication_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"publication_date"`
	Author string `json:"author"`
	Publisher string `json:"publication"`
	Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})

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