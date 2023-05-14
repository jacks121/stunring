// api/v1/header.go
package v1

import (
	"fmt"
	"net/http"
	"swetelove/services"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
)

// GetHeaderData 获取头部数据，包括所有分类和货币
func GetHeaderData(c *gin.Context) {
	db, redisClient, err := getDatabaseAndRedis(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryService, err := services.NewCategoryService(db, redisClient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category service"})
		return
	}

	currencyService, err := services.NewCurrencyService(db, redisClient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create currency service"})
		return
	}

	categories, err := categoryService.GetAllCategories()
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch categories", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	currencies, err := currencyService.GetAllCurrencies()
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch currencies", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	fmt.Println(currencies)
	headerData := map[string]interface{}{
		"categories": categories,
		"currencies": currencies,
	}

	response := utils.NewApiResponse("success", "Header data fetched successfully", headerData)
	c.JSON(http.StatusOK, response)
}
