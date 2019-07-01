package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId     string
	TotalAmount string
	CreatedBy   int
}
