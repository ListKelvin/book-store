package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	// Id 				uint64  `gorm:"primaryKey;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`
	Value  	uint16 	`gorm:"number;not null;" json:"value"`
	OrderId uint
	CustomerId uint
	Book []*Book `gorm:"many2many:books_discounts"`

}