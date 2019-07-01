package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/berthojoris/cart-backend/app/repositories/interface"
	"github.com/jinzhu/gorm"
)

type OrderServiceImpl struct {
	ItemRepository _interface.IItemRepository
}

func NewOrderServiceImpl(ItemRepository _interface.IItemRepository) *OrderServiceImpl {
	return &OrderServiceImpl{ItemRepository: ItemRepository}
}

func (s *OrderServiceImpl) GetAll(db *gorm.DB, entities interface{}) error {
	return s.ItemRepository.FindAll(db, entities.(*[]models.Children))
}

func (s *OrderServiceImpl) GetById(db *gorm.DB, entity interface{}, id int) error {
	return s.ItemRepository.FindById(db, entity.(*models.Children), id)
}

func (s *OrderServiceImpl) Create(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Create(db, entity.(*models.Children))
}

func (s *OrderServiceImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return s.ItemRepository.NewRecord(db, entity.(models.Children))
}

func (s *OrderServiceImpl) Update(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Update(db, entity.(*models.Children))
}

func (s *OrderServiceImpl) Delete(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Delete(db, entity.(*models.Children))
}
