package models

import (
	"gorm.io/gorm"
)

type Shipper struct {
	Id      uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"text;not null;" json:"name"`
	OrderId uint
}

func (s *Shipper) Prepare() {


}
func (s *Shipper) CreateShipper(db *gorm.DB) (*Shipper, error) {

	result := db.Create(&s)
	if result.Error != nil {
		return nil, result.Error // Return error if creation failed
	}
	return s, nil // Return the created book and a nil error if successful
}

func GetAllShippers(db *gorm.DB) ([]Shipper, error) {
	Shippers := make([]Shipper, 0)
	result := db.Find(&Shippers)
	if result.Error != nil {
		return nil, result.Error
	}
	return Shippers, nil
}

func (s *Shipper) GetShipperById(Id int64, db *gorm.DB) (*Shipper, error) {

	result := db.Debug().Where("ID=?", Id).Find(&s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

func (s *Shipper) DeleteShipper(ID int64, db *gorm.DB) (*Shipper, error) {

	result := db.Where("ID=?", ID).Delete(&s)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Shipper{}, nil
}