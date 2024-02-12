package models

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount   uint64 `gorm:"amount" json:"amount"`
	OrderId  uint
	BookIsbn string
}

type OrderDetailRepository interface {
	Prepare()
	CreateOrderDetail(db *gorm.DB) (*OrderDetail, error)
	GetAllOrderDetails(db *gorm.DB) ([]OrderDetail, error)
	GetOrderDetailById(Id int64, db *gorm.DB ) (*OrderDetail, error)
	DeleteOrderDetail(ID int64, db *gorm.DB ) (*OrderDetail, error)
}

func (o *OrderDetail) Prepare() {

}
func (od *OrderDetail) CreateOrderDetail(db *gorm.DB) (*OrderDetail, error) {

	result := db.Create(&od)
	if result.Error != nil {
		return nil, result.Error // Return error if creation failed
	}
	return od, nil // Return the created book and a nil error if successful
}

func GetAllOrderDetails(db *gorm.DB) ([]OrderDetail, error) {
	OrderDetails := make([]OrderDetail, 0)
	result := db.Find(&OrderDetails)
	if result.Error != nil {
		return nil, result.Error
	}
	return OrderDetails, nil
}

func (od *OrderDetail) GetOrderDetailById(Id int64, db *gorm.DB) (*OrderDetail, error) {

	result := db.Debug().Where("ID=?", Id).Find(&od)
	if result.Error != nil {
		return nil, result.Error
	}
	return od, nil
}

func (od *OrderDetail) DeleteOrderDetail(ID int64, db *gorm.DB) (*OrderDetail, error) {

	result := db.Where("ID=?", ID).Delete(&od)
	if result.Error != nil {
		return nil, result.Error
	}
	return &OrderDetail{}, nil
}