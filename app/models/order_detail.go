package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderDetail struct {
	gorm.Model
	DetailOrderid string
	ItemId int
	Qty int
}
