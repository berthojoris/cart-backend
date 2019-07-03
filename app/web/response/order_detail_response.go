package response

import (
	"github.com/berthojoris/cart-backend/app/dto/response"
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/jinzhu/gorm"
)

type OrderDetailResponse struct {
	Db *gorm.DB
}

func NewOrderDetailResponse(db *gorm.DB) OrderDetailResponse {
	return OrderDetailResponse{Db: db}
}

func (r *OrderDetailResponse) New(orderDetail models.OrderDetail) response.OrderDetail {
	response := response.OrderDetail{
		ID:      orderDetail.ID,
		OrderId: orderDetail.OrderId,
		ItemId:  orderDetail.ItemId,
		Qty:     orderDetail.Qty,
	}

	return response
}

func (r *OrderDetailResponse) Collection(orderdetails []models.OrderDetail) []response.OrderDetail {
	var responses []response.OrderDetail

	for _, orderdetail := range orderdetails {
		responses = append(responses, r.New(orderdetail))
	}

	return responses
}
