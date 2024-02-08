package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primaryKey;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`

}



type RegisterUserInput struct {
	Username     string `json:"name" binding:"required"`
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
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