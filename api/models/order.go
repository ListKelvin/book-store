package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Id uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	//customer_id
	//discount_id

	OrderDetail []OrderDetail
	Status uint16 `gorm:"boolean" json:"status"`
	// Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	Shipper Shipper
	Discount Discount

//m - m
	Customer []*Customer `gorm:"many2many:customers_orders"`
	
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