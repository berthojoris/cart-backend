package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	_interface "github.com/berthojoris/cart-backend/app/repositories/interface"
	"github.com/jinzhu/gorm"
)

type OrderDetailServiceImpl struct {
	OrderDetailRepository _interface.IOrderDetailRepository
}

func NewOrderDetailServiceImpl(OrderDetailRepository _interface.IOrderDetailRepository) *OrderDetailServiceImpl {
	return &OrderDetailServiceImpl{OrderDetailRepository: OrderDetailRepository}
}

func (s *OrderDetailServiceImpl) GetByOrderId(db *gorm.DB, entities interface{}, Id uint) error {
	return s.OrderDetailRepository.FindByOrderId(db, entities.(*[]models.OrderDetail), Id)
}

func (s *OrderDetailServiceImpl) GetAll(db *gorm.DB, entities interface{}) error {
	return s.OrderDetailRepository.FindAll(db, entities.(*[]models.OrderDetail))
}

func (s *OrderDetailServiceImpl) GetById(db *gorm.DB, entity interface{}, id int) error {
	return s.OrderDetailRepository.FindById(db, entity.(*models.OrderDetail), id)
}

func (s *OrderDetailServiceImpl) Create(db *gorm.DB, entity interface{}) error {
	return s.OrderDetailRepository.Create(db, entity.(*models.OrderDetail))
}

func (s *OrderDetailServiceImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return s.OrderDetailRepository.NewRecord(db, entity.(models.OrderDetail))
}

func (s *OrderDetailServiceImpl) Update(db *gorm.DB, entity interface{}) error {
	return s.OrderDetailRepository.Update(db, entity.(*models.OrderDetail))
}

func (s *OrderDetailServiceImpl) Delete(db *gorm.DB, entity interface{}) error {
	return s.OrderDetailRepository.Delete(db, entity.(*models.OrderDetail))
}
