package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type GenreService interface {
	Prepare()
	CreateGenre(db *gorm.DB) (models.Genre, error)
	GetAllGenres(db *gorm.DB) ([]models.Genre, error)
	GetGenreById(Id int64, db *gorm.DB) (models.Genre, error)
	DeleteGenre(ID int64, db *gorm.DB) (models.Genre, error)
}

type genreService struct {
	GenreRepository models.GenreRepository
}

// NewGenreService -> returns new Genre service
func NewGenreService(r models.GenreRepository) genreService {
	return genreService{
		GenreRepository: r,
	}
}

func (u genreService) Create(db *gorm.DB) (*models.Genre, error) {

	return u.GenreRepository.CreateGenre(db)
}

func (u genreService) GetAll(db *gorm.DB) ([]models.Genre, error) {

	return u.GenreRepository.GetAllGenres(db)
}

func (u genreService) GetById(Id int64, db *gorm.DB) (*models.Genre, error) {

	return u.GenreRepository.GetGenreById(Id, db)
}

func (u genreService) Delete(ID int64, db *gorm.DB) (*models.Genre, error) {

	return u.GenreRepository.DeleteGenre(ID, db)
}