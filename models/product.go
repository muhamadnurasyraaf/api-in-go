package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);unique"`
	Price float64
}
