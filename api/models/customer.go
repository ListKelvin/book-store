package models

import (
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

func (c *Customer) Prepare() {

	
}
func (c *Customer) CreateCustomer(db *gorm.DB) (*Customer, error) {
    
	result := db.Create(&c)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return c, nil // Return the created book and a nil error if successful
}

func  GetAllCustomers(db *gorm.DB) ([]Customer, error) {
	Customers :=  make([]Customer, 0)
	result := db.Find(&Customers)
	if result.Error != nil{
		return nil, result.Error
	}
	return Customers, nil
}

func(c *Customer) GetCustomerById(Id int64, db *gorm.DB ) (*Customer, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&c)
		if result.Error != nil{
		return nil, result.Error
	}
	return c, nil
}


func(c *Customer) DeleteCustomer(ID int64, db *gorm.DB ) (*Customer, error) {


	result:= db.Where("ID=?", ID).Delete(&c)
	if result.Error != nil{
		return nil, result.Error
	}
	return &Customer{}, nil
}