package models

import "gorm.io/gorm"

type publisher struct {
	gorm.Model
	Id 				uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`

}