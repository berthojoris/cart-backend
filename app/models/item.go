package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	gorm.Model
	ItemName string
	ItemDescription string
}
