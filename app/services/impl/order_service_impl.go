package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	_interface "github.com/berthojoris/cart-backend/app/repositories/interface"
	"github.com/jinzhu/gorm"
)

type OrderServiceImpl struct {
	OrderRepository _interface.IOrderRepository
}

func NewOrderServiceImpl(OrderRepository _interface.IOrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{OrderRepository: OrderRepository}
}

func (s *OrderServiceImpl) RemoveByOrderId(db *gorm.DB, entities interface{}, Id uint) error {
	return s.OrderRepository.DeleteByOrderId(db, entities.(*models.Order), Id)
}

func (s *OrderServiceImpl) GetAll(db *gorm.DB, entities interface{}) error {
	return s.OrderRepository.FindAll(db, entities.(*[]models.Order))
}

func (s *OrderServiceImpl) GetById(db *gorm.DB, entity interface{}, id int) error {
	return s.OrderRepository.FindById(db, entity.(*models.Order), id)
}

func (s *OrderServiceImpl) GetByOrderId(db *gorm.DB, entity interface{}, id int) error {
	return s.OrderRepository.FindById(db, entity.(*models.Order), id)
}

func (s *OrderServiceImpl) Create(db *gorm.DB, entity interface{}) error {
	return s.OrderRepository.Create(db, entity.(*models.Order))
}

func (s *OrderServiceImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return s.OrderRepository.NewRecord(db, entity.(models.Order))
}

func (s *OrderServiceImpl) Update(db *gorm.DB, entity interface{}) error {
	return s.OrderRepository.Update(db, entity.(*models.Order))
}

func (s *OrderServiceImpl) Delete(db *gorm.DB, entity interface{}) error {
	return s.OrderRepository.Delete(db, entity.(*models.Order))
}
