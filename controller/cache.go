package controller

import (
	"net/http"
	"strconv"
	"swetelove/repositories"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func FlushCache(c *gin.Context) {
	cacheService := service.NewCacheService()
	err := cacheService.SyncProductsToES()
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetLatestProducts(c *gin.Context) {
	sizeParam := c.Param("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size <= 0 {
		// 处理无效的参数值
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
		return
	}

	productService := repositories.NewProductRepository()
	products, err := productService.GetLatestProducts(size)
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	productService := repositories.NewProductRepository()
	product, err := productService.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
