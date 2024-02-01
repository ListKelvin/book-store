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
	Status uint16 `gorm:"boolean" json:"status"`
	Created_At time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}