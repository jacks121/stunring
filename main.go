package main

import (
	"path/filepath"
	"swetelove/database"
	"swetelove/routers"
	"swetelove/utils"

	"gorm.io/gorm"
)

func main() {
	configPath := filepath.Join(".", "config.toml")
	config := utils.NewConfig(configPath)

	redisClient, err := database.NewRedisClient()
	if handleError("Failed to connect to Redis: %v", err) {
		return
	}
	defer redisClient.Close()

	db, err := database.Connect(config)
	if handleError("Error connecting to database: %v", err) {
		return
	}
	defer closeDatabase(db)

	router := routers.SetupRouter(db, redisClient, config)
	router.Run(":8080")
}

func handleError(format string, err error) bool {
	if err != nil {
		utils.LogError.Printf(format, err)
		return true
	}
	return false
}

func closeDatabase(db *gorm.DB) {
	if err := database.Close(db); err != nil {
		utils.LogError.Printf("Error closing database connection: %v", err)
	}
}
