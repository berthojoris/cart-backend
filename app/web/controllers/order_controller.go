package controllers

import (
	order_request "github.com/berthojoris/cart-backend/app/dto/request/crud"
	"github.com/berthojoris/cart-backend/app/models"
	_interface "github.com/berthojoris/cart-backend/app/services/interface"
	"github.com/berthojoris/cart-backend/app/web/response"
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

	golog.Info(formRequest)

	order.TotalAmount = formRequest.Form.TotalAmount

	if err := c.OrderService.Create(c.Db, &order); err != nil {
		tx.Rollback()
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	if formRequest.Form.OrderDetail != nil && len(formRequest.Form.OrderDetail) > 0 {
		for _, orderDetailRequest := range formRequest.Form.OrderDetail {
			var detailOrder models.OrderDetail

			detailOrder.OrderId = int(order.ID)
			detailOrder.ItemId = orderDetailRequest.ItemId
			detailOrder.Qty = orderDetailRequest.Qty

			if err := c.OrderDetailService.Create(tx, &detailOrder); err != nil {
				tx.Rollback()
				response.InternalServerErrorResponse(ctx, err)
				return
			}
		}
	}

	tx.Commit()
	response.SuccessResponse(ctx, response.OK, response.SUCCESS_SAVE_ORDER, nil)
}

func (c *OrderController) UpdateOrderByIdHandler(ctx iris.Context) {

	tx := c.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			response.InternalServerErrorResponse(ctx, r)
			return
		}
	}()

	id, _ := ctx.Params().GetUint("id")

	formRequest := order_request.NewOrderRequest(ctx, c.Db, c.OrderService, c.OrderDetailService)

	if err := ctx.ReadJSON(&formRequest.Form); err != nil {
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	var order models.Order
	var orderDetail []models.OrderDetail

	c.OrderService.GetById(c.Db, &order, int(id))

	if order == (models.Order{}) {
		response.ErrorResponse(ctx, response.UNPROCESSABLE_ENTITY, "Order doesn't exists.")
		return
	}

	c.OrderDetailService.GetByOrderId(c.Db, &orderDetail, id)

	order.TotalAmount = formRequest.Form.TotalAmount

	if err := c.OrderService.Update(c.Db, &order); err != nil {
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	c.OrderDetailService.RemoveByOrderId(c.Db, &orderDetail, id)

	orderDetailResponse := response.NewOrderDetailResponse(c.Db)
	result := orderDetailResponse.Collection(orderDetail)

	tx.Commit()
	response.SuccessResponse(ctx, response.OK, response.SUCCESS_UPDATE_ORDER, result)
}

func (c *OrderController) GetOrderDetailByIdHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	var orderDetail []models.OrderDetail

	c.OrderDetailService.GetByOrderId(c.Db, &orderDetail, id)

	orderDetailResponse := response.NewOrderDetailResponse(c.Db)
	result := orderDetailResponse.Collection(orderDetail)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
}

func (c *OrderController) GetOrderByIdHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	var order models.Order

	c.OrderService.GetById(c.Db, &order, int(id))

	if order == (models.Order{}) {
		response.ErrorResponse(ctx, response.UNPROCESSABLE_ENTITY, "Order doesn't exists.")
		return
	}

	orderResponse := response.NewOrderResponse(c.Db)
	result := orderResponse.New(order)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
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
