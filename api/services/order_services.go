package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type OrderService interface {
	Prepare()
	CreateOrder(db *gorm.DB) (models.Order, error)
	GetAllOrders(db *gorm.DB) ([]models.Order, error)
	GetOrderById(Id int64, db *gorm.DB) (models.Order, error)
	DeleteOrder(ID int64, db *gorm.DB) (models.Order, error)
}

type orderService struct {
	OrderRepository models.OrderRepository
}

// NewOrderService -> returns new Order service
func NewOrderService(r models.OrderRepository) orderService {
	return orderService{
		OrderRepository: r,
	}
}

func (u orderService) Create(db *gorm.DB) (*models.Order, error) {

	return u.OrderRepository.CreateOrder(db)
}

func (u orderService) GetAll(db *gorm.DB) ([]models.Order, error) {

	return u.OrderRepository.GetAllOrders(db)
}

func (u orderService) GetById(Id int64, db *gorm.DB) (*models.Order, error) {

	return u.OrderRepository.GetOrderById(Id, db)
}

func (u orderService) Delete(ID int64, db *gorm.DB) (*models.Order, error) {

	return u.OrderRepository.DeleteOrder(ID, db)
}