package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/berthojoris/cart-backend/app/repositories/interface"
	"github.com/jinzhu/gorm"
)

type ItemServiceImpl struct {
	ItemRepository _interface.IItemRepository
}

func NewItemServiceImpl(ItemRepository _interface.IItemRepository) *ItemServiceImpl {
	return &ItemServiceImpl{ItemRepository: ItemRepository}
}

func (s *ItemServiceImpl) GetAll(db *gorm.DB, entities interface{}) error {
	return s.ItemRepository.FindAll(db, entities.(*[]models.Item))
}

func (s *ItemServiceImpl) GetById(db *gorm.DB, entity interface{}, id int) error {
	return s.ItemRepository.FindById(db, entity.(*models.Item), id)
}

func (s *ItemServiceImpl) Create(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Create(db, entity.(*models.Item))
}

func (s *ItemServiceImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return s.ItemRepository.NewRecord(db, entity.(models.Item))
}

func (s *ItemServiceImpl) Update(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Update(db, entity.(*models.Item))
}

func (s *ItemServiceImpl) Delete(db *gorm.DB, entity interface{}) error {
	return s.ItemRepository.Delete(db, entity.(*models.Item))
}
