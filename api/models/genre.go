package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	// Id 				uint64  `gorm:"primaryKey;auto_increment" json:"id"`
	Name  	string 	`gorm:"text;not null;" json:"name"`
	Book []*Book `gorm:"many2many:books_genres"`

}
type GenreRepository interface {
	Prepare()
	CreateGenre(db *gorm.DB) (*Genre, error)
	GetAllGenres(db *gorm.DB) ([]Genre, error)
	GetGenreById(Id int64, db *gorm.DB ) (*Genre, error)
	DeleteGenre(ID int64, db *gorm.DB ) (*Genre, error)
}
func (g *Genre) Prepare() {

	
}
func (g *Genre) CreateGenre(db *gorm.DB) (*Genre, error) {
    
	result := db.Create(&g)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return g, nil // Return the created book and a nil error if successful
}

func  GetAllGenres(db *gorm.DB) ([]Genre, error) {
	Genres :=  make([]Genre, 0)
	result := db.Find(&Genres)
	if result.Error != nil{
		return nil, result.Error
	}
	return Genres, nil
}

func(g *Genre) GetGenreById(Id int64, db *gorm.DB ) (*Genre, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&g)
		if result.Error != nil{
		return nil, result.Error
	}
	return g, nil
}


func(b *Genre) DeleteGenre(ID int64, db *gorm.DB ) (*Genre, error) {


	result:= db.Where("ID=?", ID).Delete(&b)
	if result.Error != nil{
		return nil, result.Error
	}
	return &Genre{}, nil
}