package _interface

import (
	"github.com/berthojoris/cart-backend/app/repositories"
)

type IItemRepository interface {
	repositories.BaseRepository
}
