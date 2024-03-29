package _interface

import (
	"github.com/berthojoris/cart-backend/app/repositories"
	"github.com/jinzhu/gorm"
)

type IOrderRepository interface {
	repositories.BaseRepository
	DeleteByOrderId(db *gorm.DB, entities interface{}, Id uint) error
}
