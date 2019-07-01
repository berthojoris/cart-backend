package controllers

import (
	order_request "github.com/berthojoris/cart-backend/app/dto/request/crud"
	"github.com/berthojoris/cart-backend/app/models"
	_interface "github.com/berthojoris/cart-backend/app/services/interface"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

type OrderController struct {
	Db                 *gorm.DB
	OrderService       _interface.IOrderService
	OrderDetailService _interface.IOrderDetailService
}

func NewOrderController(
	db *gorm.DB,
	orderService _interface.IOrderService,
	orderDetailService _interface.IOrderDetailService) *OrderController {
	return &OrderController{
		Db:                 db,
		OrderService:       orderService,
		OrderDetailService: orderDetailService,
	}
}

func (c *OrderController) SaveOrderHandler(ctx iris.Context) {
	formRequest := order_request.NewOrderRequest(ctx, c.Db, c.OrderService)

	if err := ctx.ReadJSON(&formRequest.Form); err != nil {
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	if !formRequest.Validate() {
		return
	}

	var order models.Order

	order.Orderid = formRequest.Form.Orderid
	order.TotalAmount = formRequest.Form.TotalAmount

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, nil)
}

func (c *OrderController) GetOrderHandler(ctx iris.Context) {
	var orders []models.Order

	c.OrderService.GetAll(c.Db, &orders)

	if len(orders) == 0 {
		response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, make([]interface{}, 0))
		return
	}

	orderResponse := response.NewOrderResponse(c.Db)
	result := orderResponse.Collection(orders)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
}
