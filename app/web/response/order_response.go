package response

import (
	"github.com/berthojoris/cart-backend/app/dto/response"
	"github.com/berthojoris/cart-backend/app/models"
	copier "github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type OrderResponse struct {
	Db *gorm.DB
}

func NewOrderResponse(db *gorm.DB) OrderResponse {
	return OrderResponse{Db: db}
}

func (r *OrderResponse) New(order models.Order) response.Order {

	orderDetailResponse := NewOrderDetailResponse(r.Db)

	var response response.Order
	copier.Copy(&response, &order)

	var orderDetail models.OrderDetail

	r.Db.Model(&order).Related(&orderDetail)

	if orderDetail != (models.OrderDetail{}) {
		orderDetailResponse2 := orderDetailResponse.New(orderDetail)
		response.OrderDetail = orderDetailResponse2
	}

	// response := response.Order{
	// 	ID:          order.ID,
	// 	TotalAmount: order.TotalAmount,
	// }

	return response
}

func (r *OrderResponse) Collection(orders []models.Order) []response.Order {
	var responses []response.Order

	for _, order := range orders {
		responses = append(responses, r.New(order))
	}

	return responses
}
