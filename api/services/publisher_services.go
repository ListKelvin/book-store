package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type PublisherService interface {
	Prepare()
	CreatePublisher(db *gorm.DB) (models.Publisher, error)
	GetAllPublishers(db *gorm.DB) ([]models.Publisher, error)
	GetPublisherById(Id int64, db *gorm.DB) (models.Publisher, error)
	DeletePublisher(ID int64, db *gorm.DB) (models.Publisher, error)
}

type publisherService struct {
	PublisherRepository models.PublisherRepository
}

// NewPublisherService -> returns new Publisher service
func NewPublisherService(r models.PublisherRepository) publisherService {
	return publisherService{
		PublisherRepository: r,
	}
}

func (u publisherService) Create(db *gorm.DB) (*models.Publisher, error) {

	return u.PublisherRepository.CreatePublisher(db)
}

func (u publisherService) GetAll(db *gorm.DB) ([]models.Publisher, error) {

	return u.PublisherRepository.GetAllPublishers(db)
}

func (u publisherService) GetById(Id int64, db *gorm.DB) (*models.Publisher, error) {

	return u.PublisherRepository.GetPublisherById(Id, db)
}

func (u publisherService) Delete(ID int64, db *gorm.DB) (*models.Publisher, error) {

	return u.PublisherRepository.DeletePublisher(ID, db)
}