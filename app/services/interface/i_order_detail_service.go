package _interface

import (
	"github.com/berthojoris/cart-backend/app/services"
	"github.com/jinzhu/gorm"
)

type IOrderDetailService interface {
	services.BaseService
	GetByOrderId(db *gorm.DB, entities interface{}, Id uint) error
}
