package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	// Id 				uint64  `gorm:"primaryKey;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`
	Book []*Book `gorm:"many2many:books_genres"`

}