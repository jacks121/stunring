package main

import (
	"swetelove/database"
	"swetelove/routers"
	"swetelove/utils"
)

func main() {
	database.Init()
	defer database.CloseConnections()

	router := routers.SetupRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		utils.LogError.Printf("Failed to start server: %v", err)
	}
}
