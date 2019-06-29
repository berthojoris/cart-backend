package _interface

import (
	"github.com/berthojoris/cart-backend/app/repositories"
	"github.com/jinzhu/gorm"
)

type IItemRepository interface {
	repositories.BaseRepository
}
