package models

import (
	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	gorm.Model
	DetailOrderid int
	ItemId        int
	Qty           int
}
