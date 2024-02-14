package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type DiscountService interface {
	Prepare()
	CreateDiscount(db *gorm.DB) (models.Discount, error)
	GetAllDiscounts(db *gorm.DB) ([]models.Discount, error)
	GetDiscountById(Id int64, db *gorm.DB) (models.Discount, error)
	DeleteDiscount(ID int64, db *gorm.DB) (models.Discount, error)
}

type discountService struct {
	DiscountRepository models.DiscountRepository
}

// NewDiscountService -> returns new Discount service
func NewDiscountService(r models.DiscountRepository) discountService {
	return discountService{
		DiscountRepository: r,
	}
}

func (u discountService) Create(db *gorm.DB) (*models.Discount, error) {

	return u.DiscountRepository.CreateDiscount(db)
}

func (u discountService) GetAll(db *gorm.DB) ([]models.Discount, error) {

	return u.DiscountRepository.GetAllDiscounts(db)
}

func (u discountService) GetById(Id int64, db *gorm.DB) (*models.Discount, error) {

	return u.DiscountRepository.GetDiscountById(Id, db)
}

func (u discountService) Delete(ID int64, db *gorm.DB) (*models.Discount, error) {

	return u.DiscountRepository.DeleteDiscount(ID, db)
}