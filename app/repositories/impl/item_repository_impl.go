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
	return db.Find(entities.(*[]models.Children)).Error
}

func (r *ItemRepositoryImpl) FindById(db *gorm.DB, entity interface{}, id int) error {
	return db.First(entity.(*models.Children), id).Error
}

func (r *ItemRepositoryImpl) Create(db *gorm.DB, entity interface{}) error {
	return db.Create(entity.(*models.Children)).Error
}

func (r *ItemRepositoryImpl) NewRecord(db *gorm.DB, entity interface{}) bool {
	return db.NewRecord(entity.(models.Children))
}

func (r *ItemRepositoryImpl) Update(db *gorm.DB, entity interface{}) error {
	return db.Save(entity.(*models.Children)).Error
}

func (r *ItemRepositoryImpl) Delete(db *gorm.DB, entity interface{}) error {
	return db.Delete(entity.(*models.Children)).Error
}