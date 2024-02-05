package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	//customer_id
	//discount_id

	OrderDetail []OrderDetail
	Status uint16 `gorm:"boolean" json:"status"`
	Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	Shipper Shipper
	Discount Discount

//m - m
	Customer []*Customer `gorm:"many2many:customers_orders"`
	
}