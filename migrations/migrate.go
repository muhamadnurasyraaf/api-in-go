package migrations

import (
	"log"
	"webapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(dsn string) error {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	log.Println("Database connected successfully")

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err = db.AutoMigrate(&models.Product{}); err != nil {
		return err
	}
	return nil
}
