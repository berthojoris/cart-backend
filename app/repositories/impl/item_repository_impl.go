package impl

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type ItemRepositoryImpl struct {}

func NewItemRepositoryImpl() *ItemRepositoryImpl {
	return &ItemRepositoryImpl{}
}

func (r *ItemRepositoryImpl) FindAll(db *gorm.DB, entities interface{}) error {
	return db.Find(entities.(*[]models.Item)).Error
}

func (r *ItemRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.Item), id).Error
}

func (r *ItemRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.Item)).Error
}

func (r *ItemRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.Item))
}

func (r *ItemRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.Item)).Error
}

func (r *ItemRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.Item)).Error
}