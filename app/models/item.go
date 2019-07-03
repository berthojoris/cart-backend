package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ItemName        string
	ItemDescription string
	Image           string
	Type            string
	Price           int
}
