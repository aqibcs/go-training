package models

import (
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	Name string
	Age  int
	City string
}
