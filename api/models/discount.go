package models

import (
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	// Id 				uint64  `gorm:"primaryKey;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`
	Value  	uint16 	`gorm:"number;not null;" json:"value"`
	OrderId uint
	CustomerId uint
	Book []*Book `gorm:"many2many:books_discounts"`

}
type DiscountRepository interface {
	Prepare()
	CreateDiscount(db *gorm.DB) (*Discount, error)
	GetAllDiscounts(db *gorm.DB) ([]Discount, error)
	GetDiscountById(Id int64, db *gorm.DB ) (*Discount, error)
	DeleteDiscount(ID int64, db *gorm.DB ) (*Discount, error)
}
func (d *Discount) Prepare() {

	
}
func (d *Discount) CreateDiscount(db *gorm.DB) (*Discount, error) {
    
	result := db.Create(&d)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return d, nil // Return the created book and a nil error if successful
}

func  GetAllDiscounts(db *gorm.DB) ([]Discount, error) {
	Discounts :=  make([]Discount, 0)
	result := db.Find(&Discounts)
	if result.Error != nil{
		return nil, result.Error
	}
	return Discounts, nil
}

func(d *Discount) GetDiscountById(Id int64, db *gorm.DB ) (*Discount, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&d)
		if result.Error != nil{
		return nil, result.Error
	}
	return d, nil
}


func(d *Book) DeleteDiscount(ID int64, db *gorm.DB ) (*Discount, error) {


	result:= db.Where("ID=?", ID).Delete(&d)
	if result.Error != nil{
		return nil, result.Error
	}
	return &Discount{}, nil
}