package models

import "gorm.io/gorm"

type genre struct {
	gorm.Model
	Id 				uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`

}