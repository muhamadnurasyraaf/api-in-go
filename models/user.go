package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;size:255"`
	Password string
}
