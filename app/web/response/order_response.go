package response

import (
	"github.com/berthojoris/cart-backend/app/dto/response"
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type OrderResponse struct {
	Db *gorm.DB
}

func NewOrderResponse(db *gorm.DB) OrderResponse {
	return OrderResponse{Db: db}
}

func (r *OrderResponse) New(order models.Order) response.Order {
	response := response.Order{
		ID:          order.ID,
		OrderId:     order.OrderId,
		TotalAmount: order.TotalAmount,
		CreatedBy:   order.CreatedBy,
	}

	return response
}

func (r *OrderResponse) Collection(orders []models.Order) []response.Order {
	var responses []response.Order

	for _, order := range orders {
		responses = append(responses, r.New(order))
	}

	return responses
}
