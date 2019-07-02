package controllers

import (
	order_request "github.com/berthojoris/cart-backend/app/dto/request/crud"
	"github.com/berthojoris/cart-backend/app/models"
	_interface "github.com/berthojoris/cart-backend/app/services/interface"
	"github.com/berthojoris/cart-backend/app/web/response"
	copier "github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	golog "github.com/kataras/golog"
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

	tx := c.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			response.InternalServerErrorResponse(ctx, r)
			return
		}
	}()

	formRequest := order_request.NewOrderRequest(ctx, c.Db, c.OrderService, c.OrderDetailService)

	if err := ctx.ReadJSON(&formRequest.Form); err != nil {
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	if !formRequest.Validate() {
		return
	}

	var order models.Order

	golog.Info(order)

	order.OrderId = formRequest.Form.OrderId
	order.TotalAmount = formRequest.Form.TotalAmount

	if err := c.OrderService.Create(c.Db, &order); err != nil {
		tx.Rollback()
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	if formRequest.Form.OrderDetail != nil && len(formRequest.Form.OrderDetail) > 0 {
		for _, orderDetailRequest := range formRequest.Form.OrderDetail {
			var detailOrder models.OrderDetail

			copier.Copy(&detailOrder, &orderDetailRequest)

			detailOrder.DetailOrderid = int(formRequest.Form.OrderId)

			if err := c.OrderDetailService.Create(tx, &detailOrder); err != nil {
				tx.Rollback()
				response.InternalServerErrorResponse(ctx, err)
				return
			}
		}
	}

	golog.Info(order)

	tx.Commit()
	response.SuccessResponse(ctx, response.OK, response.SUCCESS_SAVE_ORDER, nil)
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
