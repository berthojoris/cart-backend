package _interface

import (
	"github.com/berthojoris/cart-backend/app/repositories"
	"github.com/jinzhu/gorm"
)

type IOrderDetailRepository interface {
	repositories.BaseRepository
	FindByOrderId(db *gorm.DB, entities interface{}, Id uint) error
	DeleteByOrderId(db *gorm.DB, entities interface{}, Id uint) error
}
