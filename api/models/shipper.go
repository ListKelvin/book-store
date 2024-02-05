package models

type Shipper struct {
	Id      uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"text;not null;" json:"name"`
	OrderId uint16
}