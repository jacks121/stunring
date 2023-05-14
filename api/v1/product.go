package v1

import (
	"errors"
	"net/http"
	"strconv"
	"swetelove/services"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func getProductService(c *gin.Context) (*services.ProductService, error) {
	db, ok := c.Get("db")
	if !ok {
		return nil, errors.New("database not found in context")
	}

	redisClient, ok := c.Get("redis")
	if !ok {
		return nil, errors.New("redis client not found in context")
	}

	productService, err := services.NewProductService(db.(*gorm.DB), redisClient.(*redis.Client))
	if err != nil {
		return nil, err
	}

	return productService, nil
}

func GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	productService, err := getProductService(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product, err := productService.GetProductByID(uint(productID))
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch product", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	response := utils.NewApiResponse("success", "Product fetched successfully", product)
	c.JSON(http.StatusOK, response)
}
