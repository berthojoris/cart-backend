package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Orderid     string
	TotalAmount string
	CreatedBy   int
}
