package response

import (
	"github.com/berthojoris/cart-backend/app/dto/response"
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type ItemResponse struct {
	Db *gorm.DB
}

func NewItemResponse(db *gorm.DB) ItemResponse {
	return ItemResponse{Db: db}
}

func (r *ItemResponse) New(item models.Item) response.Item {
	response := response.Item{
		ID:              item.ID,
		ItemName:        item.ItemName,
		ItemDescription: item.ItemDescription,
		Image:           item.Image,
		Type:            item.Type,
		Price:           item.Price,
	}

	return response
}

func (r *ItemResponse) Collection(items []models.Item) []response.Item {
	var responses []response.Item

	for _, item := range items {
		responses = append(responses, r.New(item))
	}

	return responses
}
