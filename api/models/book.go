package models

import (
	"time"

	// "github.com/ListKelvin/book-store/pkg/models"
	"gorm.io/gorm"
)

// var db *gorm.DB

type Book struct {
	// gorm.Model
	Isbn string `gorm:"primaryKey;size:13;unique_index" json:"isbn"`
	Title string `gorm:"size:255;not null;unique" json:"title"`
	Available_Quantity uint16 `gorm:"default:0" json:"available_quantity"`
	Price uint16 `gorm:"default:0" json:"price,omitempty"`
	Edition uint8 `gorm:"default:0" json:"edition"`
	Publication_Date time.Time `json:"publication_date"`
// belongs to
	Author Author `json:"author"`
	AuthorID uint32 `gorm:"not null" json:"author_id"`
// belongs to
	Publisher Publisher `json:"publication"`
	PublisherID uint32 `gorm:"not null" json:"publisher_id"`

	// Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
// m2m a book has many genre
	Genre []*Genre `gorm:"many2many:books_genres;"`
	Discount []*Discount `gorm:"many2many:books_discounts;"`


// o - m 
	OrderDetail []OrderDetail
	Review []Review

}
// func init() {

// 	db = config.GetDB()


// }

func (b *Book) CreateBook(db *gorm.DB) (*Book, error) {
    
	result := db.Create(&b)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return b, nil // Return the created book and a nil error if successful
}

func GetAllBooks(db *gorm.DB) []Book{
	Books :=  make([]Book, 0)
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64,db *gorm.DB ) (*Book, *gorm.DB) {
	var getBook Book

	result:= db.Where("ID=?", Id).Find(&getBook)
	return &getBook, result
}


func DeleteBook(ID int64, db *gorm.DB ) Book {
	var book Book

	 db.Where("ID=?", ID).Delete(book)
	return book
}