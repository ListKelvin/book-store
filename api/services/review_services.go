package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type ReviewService interface {
	Prepare()
	CreateReview(db *gorm.DB) (models.Review, error)
	GetAllReviews(db *gorm.DB) ([]models.Review, error)
	GetReviewById(Id int64, db *gorm.DB) (models.Review, error)
	DeleteReview(ID int64, db *gorm.DB) (models.Review, error)
}

type reviewService struct {
	ReviewRepository models.ReviewRepository
}

// NewReviewService -> returns new Review service
func NewReviewService(r models.ReviewRepository) reviewService {
	return reviewService{
		ReviewRepository: r,
	}
}

func (u reviewService) Create(db *gorm.DB) (*models.Review, error) {

	return u.ReviewRepository.CreateReview(db)
}

func (u reviewService) GetAll(db *gorm.DB) ([]models.Review, error) {

	return u.ReviewRepository.GetAllReviews(db)
}

func (u reviewService) GetById(Id int64, db *gorm.DB) (*models.Review, error) {

	return u.ReviewRepository.GetReviewById(Id, db)
}

func (u reviewService) Delete(ID int64, db *gorm.DB) (*models.Review, error) {

	return u.ReviewRepository.DeleteReview(ID, db)
}