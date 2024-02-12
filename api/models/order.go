package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Id uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	//customer_id
	//discount_id

	OrderDetail []OrderDetail
	Status uint16 `gorm:"boolean" json:"status"`
	// Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	Shipper Shipper
	Discount Discount

//m - m
	Customer []*Customer `gorm:"many2many:customers_orders"`
	
}

type OrderRepository interface {
	Prepare()
	CreateOrder(db *gorm.DB) (*Order, error)
	GetAllOrders(db *gorm.DB) ([]Order, error)
	GetOrderById(Id int64, db *gorm.DB ) (*Order, error)
	DeleteOrder(ID int64, db *gorm.DB ) (*Order, error)
}
func (o *Order) Prepare() {

	
}
func (o *Order) CreateOrder(db *gorm.DB) (*Order, error) {
    
	result := db.Create(&o)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return o, nil // Return the created book and a nil error if successful
}

func  GetAllOrders(db *gorm.DB) ([]Order, error) {
	Orders :=  make([]Order, 0)
	result := db.Find(&Orders)
	if result.Error != nil{
		return nil, result.Error
	}
	return Orders, nil
}

func(o *Order) GetOrderById(Id int64, db *gorm.DB ) (*Order, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&o)
		if result.Error != nil{
		return nil, result.Error
	}
	return o, nil
}


func(o *Order) DeleteOrder(ID int64, db *gorm.DB ) (*Order, error) {


	result:= db.Where("ID=?", ID).Delete(&o)
	if result.Error != nil{
		return nil, result.Error
	}
	return &Order{}, nil
}

