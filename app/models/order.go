package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	TotalAmount uint
}
