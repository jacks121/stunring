// api/v1/category.go
package v1

import (
	"net/http"
	"strconv"
	"swetelove/services"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
)

// GetCategories 获取所有分类
func GetCategories(c *gin.Context) {
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

	categories, err := categoryService.GetAllCategories()
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch categories", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(categories) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "No categories found",
		})
		return
	}

	response := utils.NewApiResponse("success", "Category list fetched successfully", categories)
	c.JSON(http.StatusOK, response)
}

// GetProductsByCategory 根据分类ID获取产品
func GetProductsByCategory(c *gin.Context) {
	db, redisClient, err := getDatabaseAndRedis(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	sortType := c.DefaultQuery("sort", "default")
	isDesc := c.DefaultQuery("desc", "false") == "true"

	categoryService, err := services.NewCategoryService(db, redisClient)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category service"})
		return
	}

	products, err := categoryService.GetProductsByCategoryID(uint(categoryID), sortType, isDesc)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch products for the given category ID", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(products) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "No products found for the given category ID",
		})
		return
	}

	response := utils.NewApiResponse("success", "Product list fetched successfully for the given category ID", products)
	c.JSON(http.StatusOK, response)
}
