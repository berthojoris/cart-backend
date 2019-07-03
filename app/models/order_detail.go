package models

import (
	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	gorm.Model
	OrderId int
	ItemId  int
	Qty     int
}
