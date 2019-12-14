package models

import (
	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model
	Source   string
	Language string
}
