package v1

import (
	"errors"
	"net/http"
	"swetelove/services"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func getDatabaseAndRedis(c *gin.Context) (*gorm.DB, *redis.Client, error) {
	db, ok := c.Get("db")
	if !ok {
		return nil, nil, errors.New("database not found in context")
	}

	redisClient, ok := c.Get("redis")
	if !ok {
		return nil, nil, errors.New("redis client not found in context")
	}

	return db.(*gorm.DB), redisClient.(*redis.Client), nil
}

func GetAdvertisementsByPosition(c *gin.Context) {
	position := c.Param("position")

	db, redisClient, err := getDatabaseAndRedis(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	advertisementService, err := services.NewAdvertisementService(db, redisClient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create advertisement service"})
		return
	}

	advertisements, err := advertisementService.GetAdvertisementsByPosition(position)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch advertisements", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(advertisements) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "No advertisements found for the given position",
		})
		return
	}

	response := utils.NewApiResponse("success", "Advertisement list fetched successfully", advertisements)
	c.JSON(http.StatusOK, response)
}
