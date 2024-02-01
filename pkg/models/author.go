package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	Id 				uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Company_Name  	string 	`gorm:"text;not null;" json:"company_name"`
	First_Name    	string 	`gorm:"text;not null;" json:"first_name"`
	Second_Name  	string 	`gorm:"text;not null;" json:"second_name"`
}