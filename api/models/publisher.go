package models

import (
	"gorm.io/gorm"
)

type Publisher struct {
	Id       uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Name     string `gorm:"text;not null;" json:"name"`
	Phone_No string `gorm:"default:text;not null" json:"phone_no"`
}

func (p *Publisher) Prepare() {


}
func (b *Publisher) CreatePublisher(db *gorm.DB) (*Publisher, error) {

	result := db.Create(&b)
	if result.Error != nil {
		return nil, result.Error // Return error if creation failed
	}
	return b, nil // Return the created book and a nil error if successful
}

func GetAllPublishers(db *gorm.DB) ([]Publisher, error) {
	Publishers := make([]Publisher, 0)
	result := db.Find(&Publishers)
	if result.Error != nil {
		return nil, result.Error
	}
	return Publishers, nil
}

func (p *Publisher) GetPublisherById(Id int64, db *gorm.DB) (*Publisher, error) {

	result := db.Debug().Where("ID=?", Id).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

func (p *Publisher) DeletePublisher(ID int64, db *gorm.DB) (*Publisher, error) {

	result := db.Where("ID=?", ID).Delete(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Publisher{}, nil
}