package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Age       int
	City      string
}
