package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Price float64
}
