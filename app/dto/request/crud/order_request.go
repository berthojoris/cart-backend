package crud

import (
	"github.com/berthojoris/cart-backend/app/dto/request"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type FormOrderRequest struct {
	ID *uint `json:"id"`
	Orderid string `json:"order_id" validate:"required"`
	TotalAmount string `json:"total_amount" validate:"required"`
	CreatedBy int `json:"created_by" validate:"required"`
	CreatedDate string `json:"created_date" validate:"required"`
	UpdatedDate string `json:"updated_date" validate:"required"`
}

type OrderRequest struct {
	Ctx iris.Context
	Db *gorm.DB
	Form FormOrderRequest
}

func NewOrderRequest(ctx iris.Context, db *gorm.DB) OrderRequest {
	return OrderRequest{
		Ctx: ctx,
		Db: db,
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