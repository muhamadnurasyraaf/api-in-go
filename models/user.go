package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);"`
	Password string
	Branch   []Branch `gorm:"many2many:user_branches;"`
}
