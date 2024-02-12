package models

import (
	"html"
	"strings"
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
type BookRepository interface {
	Prepare()
	CreateBook(db *gorm.DB) (*Book, error)
	GetAllBooks(db *gorm.DB) ([]Book, error)
	GetBookById(Id int64, db *gorm.DB ) (*Book, error)
	DeleteBook(ID int64, db *gorm.DB ) (*Book, error)
}
func (b *Book) Prepare() {
	b.Isbn= "";
	b.Title = html.EscapeString(strings.TrimSpace(b.Title));
	b.Available_Quantity = 0;
	b.Edition= 0
	b.Genre = []*Genre{}
	b.Price = 0
	b.Author = Author{}
	b.Discount = []*Discount{}
	b.Publication_Date = time.Now()
	b.Review = []Review{}
	
}
func (b *Book) CreateBook(db *gorm.DB) (*Book, error) {
    
	result := db.Create(&b)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return b, nil // Return the created book and a nil error if successful
}

func  GetAllBooks(db *gorm.DB) ([]Book, error) {
	Books :=  make([]Book, 0)
	result := db.Find(&Books)
	if result.Error != nil{
		return nil, result.Error
	}
	return Books, nil
}

func(b *Book) GetBookById(Id int64, db *gorm.DB ) (*Book, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&b)
		if result.Error != nil{
		return nil, result.Error
	}
	return b, nil
}


func(b *Book) DeleteBook(ID int64, db *gorm.DB ) (*Book, error) {


	result:= db.Where("ID=?", ID).Delete(&b)
	if result.Error != nil{
		return nil, result.Error
	}
	return &Book{}, nil
}