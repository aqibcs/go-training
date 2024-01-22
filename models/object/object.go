package models

import (
	"time"

	"gorm.io/gorm"
)

// Employee represents an employee in the organization.
type Employee struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      // Time when the employee record was created
	UpdatedAt time.Time      // Time when the employee record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft deletion field, automatically set when record is deleted

	Name string `gorm:"not null;index"` // Employee's name (indexed and required)
	Age  int    `gorm:"not null"`       // Employee's age (required)
	City string `gorm:"not null;index"` // Employee's city (indexed and required)
}
