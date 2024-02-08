package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null;unique" json:"first_name"`
	LastName string `gorm:"size:255;not null;unique" json:"last_name"`
	FlatNo uint16 `gorm:"default:0" json:"flat_no,omitempty"`
	PhoneNumber uint16 `gorm:"unique;not null" json:"phone_number"`
	City string `gorm:"size:255" json:"city"`
	BuildingNo string `gorm:"size:255" json:"building_no"`
	Street string `gorm:"size:255" json:"street"`
	PostalCode string `gorm:"size:255,unique" json:"postal_code"`
	PasswordHash string `gorm:"not null" json:"password_hash"`

//1-m
	Discount []Discount
	Review []Review

// m - m
	Order []*Order `gorm:"many2many:customers_orders;"`
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