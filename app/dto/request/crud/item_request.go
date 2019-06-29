package crud

import (
	"github.com/berthojoris/cart-backend/app/dto/request"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type FormItemRequest struct {
	ID *uint `json:"id"`
	ItemName string `json:"item_name" validate:"required"`
	ItemDescription string `json:"item_description" validate:"required"`
}

type ItemRequest struct {
	Ctx iris.Context
	Db *gorm.DB
	Form FormItemRequest
}

func NewItemRequest(ctx iris.Context, db *gorm.DB) ItemRequest {
	return ItemRequest{
		Ctx: ctx,
		Db: db,
	}
}

func (r *ItemRequest) Validate() bool {
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
