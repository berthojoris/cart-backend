package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type OrderDetailRepositoryImpl struct {}

func NewOrderDetailRepositoryImpl() *OrderDetailRepositoryImpl {
	return &OrderDetailRepositoryImpl{}
}

func (r *OrderDetailRepositoryImpl) FindAll(db *gorm.DB, entities interface{}) error {
	return db.Find(entities.(*[]models.Children)).Error
}

func (r *OrderDetailRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.Children), id).Error
}

func (r *OrderDetailRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.Children)).Error
}

func (r *OrderDetailRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.Children))
}

func (r *OrderDetailRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.Children)).Error
}

func (r *OrderDetailRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.Children)).Error
}