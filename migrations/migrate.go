package migrations

import (
	"log"
	"webapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.Product{}); err != nil {
		return nil, err
	}
	return db, nil
}
