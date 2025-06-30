package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	// ID     string `gorm:"type:uuid;primaryKey"`
	Title  string `gorm:"size:256;unique"`
	Status string
	UserId uint

	user User
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
