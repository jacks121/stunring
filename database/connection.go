// database/connection.go
package database

import (
	"fmt"
	"swetelove/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config *utils.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Get("Database.User"), config.Get("Database.Password"), config.Get("Database.Host"), config.Get("Database.Port"), config.Get("Database.Name"), config.Get("Database.Charset"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL database: %w", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}
