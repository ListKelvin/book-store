package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	Id         uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	BookIsbn   string `json:"book_isbn"`
	CustomerId uint   `json:"customer_id"`
	Review_No  uint8  `gorm:"default:0" json:"review_no"`
	// Create_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_date,omitempty"`
	// Updated_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_date,omitempty"`
}

func (b *Book) Prepare() {
	b.Isbn = ""
	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Available_Quantity = 0
	b.Edition = 0
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

func GetAllBooks(db *gorm.DB) ([]Book, error) {
	Books := make([]Book, 0)
	result := db.Find(&Books)
	if result.Error != nil {
		return nil, result.Error
	}
	return Books, nil
}

func (b *Book) GetBookById(Id int64, db *gorm.DB) (*Book, error) {

	result := db.Debug().Where("ID=?", Id).Find(&b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}

func (b *Book) DeleteBook(ID int64, db *gorm.DB) (*Book, error) {

	result := db.Where("ID=?", ID).Delete(&b)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Book{}, nil
}