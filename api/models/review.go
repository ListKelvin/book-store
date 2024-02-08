package models

type Review struct {


	Id uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	BookIsbn string `json:"book_isbn"`
	CustomerId uint `json:"customer_id"`
	Review_No uint8 `gorm:"default:0" json:"review_no"`
	// Create_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_date,omitempty"`
	// Updated_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_date,omitempty"`
}