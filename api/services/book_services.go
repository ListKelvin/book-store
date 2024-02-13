package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type BookService interface {
	Prepare()
	CreateBook(db *gorm.DB) (models.Book, error)
	GetAllBooks(db *gorm.DB) ([]models.Book, error)
	GetBookById(Id int64, db *gorm.DB) (models.Book, error)
	DeleteBook(ID int64, db *gorm.DB) (models.Book, error)
}

type bookService struct {
	BookRepository models.BookRepository
}

// NewBookService -> returns new Book service
func NewBookService(r models.BookRepository) bookService {
	return bookService{
		BookRepository: r,
	}
}

func (u bookService) Create(db *gorm.DB) (*models.Book, error) {

	return u.BookRepository.CreateBook(db)
}

func (u bookService) GetAll(db *gorm.DB) ([]models.Book, error) {

	return u.BookRepository.GetAllBooks(db)
}

func (u bookService) GetById(Id int64, db *gorm.DB) (*models.Book, error) {

	return u.BookRepository.GetBookById(Id, db)
}

func (u bookService) Delete(ID int64, db *gorm.DB) (*models.Book, error) {

	return u.BookRepository.DeleteBook(ID, db)
}