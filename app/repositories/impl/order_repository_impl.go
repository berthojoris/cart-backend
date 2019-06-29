package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type OrderRepositoryImpl struct {}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) FindAll(db *gorm.DB, entities interface{}) error {
	return db.Find(entities.(*[]models.Children)).Error
}

func (r *OrderRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.Children), id).Error
}

func (r *OrderRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.Children)).Error
}

func (r *OrderRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.Children))
}

func (r *OrderRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.Children)).Error
}

func (r *OrderRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.Children)).Error
}