package crud

import (
	"github.com/berthojoris/cart-backend/app/dto/request"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type FormOrderDetailRequest struct {
	ID            *uint `json:"id"`
	DetailOrderid int   `json:"detail_order_id" validate:"required"`
	ItemId        int   `json:"item_id" validate:"required"`
	Qty           int   `json:"qty" validate:"required"`
}

type OrderDetailRequest struct {
	Ctx  iris.Context
	Db   *gorm.DB
	Form FormOrderDetailRequest
}

func NewOrderDetailRequest(ctx iris.Context, db *gorm.DB) OrderDetailRequest {
	return OrderDetailRequest{
		Ctx: ctx,
		Db:  db,
	}
}

func (r *OrderDetailRequest) Validate() bool {
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
