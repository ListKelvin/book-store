package services

import (
	"github.com/ListKelvin/book-store/api/models"
	"gorm.io/gorm"
)

type OrderDetailService interface {
	Prepare()
	CreateOrderDetail(db *gorm.DB) (models.OrderDetail, error)
	GetAllOrderDetails(db *gorm.DB) ([]models.OrderDetail, error)
	GetOrderDetailById(Id int64, db *gorm.DB) (models.OrderDetail, error)
	DeleteOrderDetail(ID int64, db *gorm.DB) (models.OrderDetail, error)
}

type orderDetailService struct {
	OrderDetailRepository models.OrderDetailRepository
}

// NewOrderDetailService -> returns new OrderDetail service
func NewOrderDetailService(r models.OrderDetailRepository) orderDetailService {
	return orderDetailService{
		OrderDetailRepository: r,
	}
}

func (u orderDetailService) Create(db *gorm.DB) (*models.OrderDetail, error) {

	return u.OrderDetailRepository.CreateOrderDetail(db)
}

func (u orderDetailService) GetAll(db *gorm.DB) ([]models.OrderDetail, error) {

	return u.OrderDetailRepository.GetAllOrderDetails(db)
}

func (u orderDetailService) GetById(Id int64, db *gorm.DB) (*models.OrderDetail, error) {

	return u.OrderDetailRepository.GetOrderDetailById(Id, db)
}

func (u orderDetailService) Delete(ID int64, db *gorm.DB) (*models.OrderDetail, error) {

	return u.OrderDetailRepository.DeleteOrderDetail(ID, db)
}