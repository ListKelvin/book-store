package models

import "gorm.io/gorm"

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

