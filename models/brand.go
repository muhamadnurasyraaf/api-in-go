package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);unique"`
}
