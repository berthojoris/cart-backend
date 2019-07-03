package crud

import (
	"github.com/berthojoris/cart-backend/app/dto/request"
	_interface "github.com/berthojoris/cart-backend/app/services/interface"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type FormOrderRequest struct {
	ID          *uint                    `json:"id"`
	OrderId     int                   `json:"order_id" validate:"required"`
	TotalAmount uint                     `json:"total_amount" validate:"required"`
	OrderDetail []FormOrderDetailRequest `json:"detail" validate:"required,dive"`
}

type OrderRequest struct {
	Ctx                iris.Context
	Db                 *gorm.DB
	Form               FormOrderRequest
	OrderService       _interface.IOrderService
	OrderDetailService _interface.IOrderDetailService
}

func NewOrderRequest(ctx iris.Context, db *gorm.DB, orderService _interface.IOrderService, orderDetailService _interface.IOrderDetailService) OrderRequest {
	return OrderRequest{
		Ctx:                ctx,
		Db:                 db,
		OrderService:       orderService,
		OrderDetailService: orderDetailService,
	}
}

func (r *OrderRequest) Validate() bool {
	baseRequest := request.New()
	var validationErrors []string

	err := baseRequest.Validate.Struct(r.Form)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, e.Translate(baseRequest.Trans))
		}
	}

	if len(validationErrors) > 0 {
		response.ValidationResponse(r.Ctx, response.BAD_REQUEST_MESSAGE, validationErrors)
		return false
	}

	return true
}
