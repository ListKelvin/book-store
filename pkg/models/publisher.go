package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Id 				uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name  			string 	`gorm:"text;not null;" json:"name"`
	Phone_No string `gorm:"default:text;not null" json:"phone_no"`

}