package models

type Author struct {
	// gorm.Model
	Id           uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Company_Name string `gorm:"text;not null;" json:"company_name"`
	First_Name   string `gorm:"text;not null;" json:"first_name"`
	Second_Name  string `gorm:"text;not null;" json:"second_name"`
}