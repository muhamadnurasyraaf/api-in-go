package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"size:255" jason:"name" validate:"required,min=2,max=100"`
	Price float64 `json:"price" validate:"required,gt=0"`
}
