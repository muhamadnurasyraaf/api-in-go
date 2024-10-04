package models

import "gorm.io/gorm"

type Size struct {
	gorm.Model
	ID           uint
	Name         string `gorm:"type:varchar(255);"`
	Width        int
	Height       int
	Rim_Diameter int
	Load_Index   int
	Speed_Index  int
	Speed_Rating string `gorm:"type:varchar(50);"`
}
