package models

type Publisher struct {
	Id       uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Name     string `gorm:"text;not null;" json:"name"`
	Phone_No string `gorm:"default:text;not null" json:"phone_no"`
}