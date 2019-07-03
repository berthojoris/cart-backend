package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId     int
	TotalAmount uint
}
