package database

import (
	"log"
	"webapp/models"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var tables = map[string]interface{}{
	"branches": &models.Branch{},
	"sizes":    &models.Size{},
	"brands":   &models.Brand{},
	"users":    &models.User{},
	"products": &models.Product{},
}

// initDB initializes the database connection.
func initDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	log.Println("Database connected successfully")
	return db, nil
}

func autoMigrateTables(db *gorm.DB) error {
	for _, model := range tables {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

// dropTables drops existing tables.
func dropTables(db *gorm.DB) error {
	for index, model := range tables {
		if db.Migrator().HasTable(index) {
			if err := db.Migrator().DropTable(model); err != nil {
				return err
			}
			log.Printf("Table %s dropped successfully", index)
		}
	}
	return nil
}

func Migrate(dsn string) (*gorm.DB, error) {
	db, err := initDB(dsn)
	if err != nil {
		return nil, err
	}

	if err = autoMigrateTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

// MigrateFresh handles the migration process with dropping existing tables.
func MigrateFresh(dsn string) (*gorm.DB, error) {
	db, err := initDB(dsn)
	if err != nil {
		return nil, err
	}

	if err = dropTables(db); err != nil {
		return nil, err
	}

	if err = autoMigrateTables(db); err != nil {
		return nil, err
	}

	return db, nil
}
