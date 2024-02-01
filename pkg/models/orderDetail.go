package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	Id     uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount uint64 `gorm:"amount" json:"amount"`
}