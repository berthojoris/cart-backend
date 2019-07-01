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

func (r *ItemResponse) New(itemdata models.Item) response.Item {
	response := response.Item{
		ID: itemdata.ID,
		ItemName: itemdata.ItemName,
		ItemDescription: itemdata.ItemDescription,
		Image: itemdata.Image,
		Type: itemdata.Type,
	}

	return response
}

func (r *ItemResponse) Collection(items []models.Item) []response.Item {
	var responses []response.Item

	for _, itemdata := range items {
		responses = append(responses, r.New(itemdata))
	}

	return responses
}
