package database

import (
	"log"
	"webapp/models"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

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
