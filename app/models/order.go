package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Orderid string
	TotalAmount string
	CreatedBy int
	CreatedDate time.Time
	UpdatedDate time.Time
}
