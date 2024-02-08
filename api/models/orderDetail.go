package models

type OrderDetail struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount   uint64 `gorm:"amount" json:"amount"`
	OrderId  uint
	BookIsbn string
}