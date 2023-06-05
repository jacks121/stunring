package controller

import (
	"net/http"
	"strconv"
	"swetelove/repositories"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

type CacheController struct {
	CacheService *service.CacheService
}

func NewCacheController() *CacheController {
	return &CacheController{
		CacheService: service.NewCacheService(),
	}
}

func (cc *CacheController) FlushCache(c *gin.Context) {
	err := cc.CacheService.SyncProductsToES()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

type CacheProductController struct {
	ProductService *repositories.ProductRepository
}

func NewCacheProductController() *CacheProductController {
	return &CacheProductController{
		ProductService: repositories.NewProductRepository(),
	}
}

func (cpc *CacheProductController) GetLatestProducts(c *gin.Context) {
	sizeParam := c.Param("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
		return
	}

	products, err := cpc.ProductService.GetLatestProducts(size, "created_at")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (cpc *CacheProductController) GetProductByID(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := cpc.ProductService.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
