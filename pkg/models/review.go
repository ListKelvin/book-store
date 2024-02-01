package models

import (
	"time"

	"gorm.io/gorm"
)

type review struct {
	gorm.Model

	Id uint16 `gorm:"primary_key;autoIncrement" json:"id"`
	//book_id
	//customer_id
	Review_No uint8 `gorm:"default:0" json:"review_no"`
	Create_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_date,omitempty"`
	Updated_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_date,omitempty"`
}