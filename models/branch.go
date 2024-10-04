package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);unique"`
	User []User `gorm:"many2many:user_branches;"`
}
