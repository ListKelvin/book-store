package models

import "gorm.io/gorm"

type Author struct {
	// gorm.Model
	Id           uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Company_Name string `gorm:"text;not null;" json:"company_name"`
	First_Name   string `gorm:"text;not null;" json:"first_name"`
	Second_Name  string `gorm:"text;not null;" json:"second_name"`
}

type AuthorRepository interface {
	Prepare()
	CreateAuthor(db *gorm.DB) (*Author, error)
	GetAllAuthors(db *gorm.DB) ([]Author, error)
	GetAuthorById(Id int64, db *gorm.DB ) (*Author, error)
	DeleteAuthor(ID int64, db *gorm.DB ) (*Author, error)
}


func (a *Author) CreateAuthor(db *gorm.DB) (*Author, error) {

	result := db.Create(&a)
	if result.Error != nil {
		return nil, result.Error // Return error if creation failed
	}
	return a, nil // Return the created Author and a nil error if successful
}

func GetAllAuthors(db *gorm.DB) ([]Author, error) {
	Authors := make([]Author, 0)
	result := db.Find(&Authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return Authors, nil
}

func (a *Author) GetAuthorById(Id int64, db *gorm.DB) (*Author, error) {

	result := db.Debug().Where("ID=?", Id).Find(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

func (a *Author) DeleteAuthor(ID int64, db *gorm.DB) (*Author, error) {

	result := db.Where("ID=?", ID).Delete(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Author{}, nil
}