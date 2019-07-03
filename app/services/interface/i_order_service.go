package _interface

import (
	"github.com/berthojoris/cart-backend/app/services"
	"github.com/jinzhu/gorm"
)

type IOrderService interface {
	services.BaseService
	RemoveByOrderId(db *gorm.DB, entities interface{}, Id uint) error
}
