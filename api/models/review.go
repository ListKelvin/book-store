package models

import (
	"gorm.io/gorm"
)

type Review struct {
	Id         uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	BookIsbn   string `json:"book_isbn"`
	CustomerId uint   `json:"customer_id"`
	Review_No  uint8  `gorm:"default:0" json:"review_no"`
	// Create_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_date,omitempty"`
	// Updated_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_date,omitempty"`
}
type ReviewRepository interface {
	Prepare()
	CreateReview(db *gorm.DB) (*Review, error)
	GetAllReviews(db *gorm.DB) ([]Review, error)
	GetReviewById(Id int64, db *gorm.DB ) (*Review, error)
	DeleteReview(ID int64, db *gorm.DB ) (*Review, error)
}

func (r *Review) Prepare() {


}
func (r *Review) CreateReview(db *gorm.DB) (*Review, error) {

	result := db.Create(&r)
	if result.Error != nil {
		return nil, result.Error // Return error if creation failed
	}
	return r, nil // Return the created book and a nil error if successful
}

func GetAllReviews(db *gorm.DB) ([]Review, error) {
	Reviews := make([]Review, 0)
	result := db.Find(&Reviews)
	if result.Error != nil {
		return nil, result.Error
	}
	return Reviews, nil
}

func (r *Review) GetReviewById(Id int64, db *gorm.DB) (*Review, error) {

	result := db.Debug().Where("ID=?", Id).Find(&r)
	if result.Error != nil {
		return nil, result.Error
	}
	return r, nil
}

func (r *Review) DeleteReview(ID int64, db *gorm.DB) (*Review, error) {

	result := db.Where("ID=?", ID).Delete(&r)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Review{}, nil
}