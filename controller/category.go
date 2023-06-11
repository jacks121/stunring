package controller

import (
	"net/http"
	"strconv"
	"strings"
	"swetelove/repositories"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	productService *service.ProductService
}

// NewCategoryController 创建 CategoryController 实例
func NewCategoryController() *CategoryController {
	return &CategoryController{
		productService: service.NewProductService(),
	}
}

// Show 处理显示分类的请求
func (cc *CategoryController) Show(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Create a new ProductsFilter and populate it with the URL parameters
	filter := repositories.ProductsFilter{}
	if caratRange := c.Query("carat_range"); caratRange != "" {
		filter.CaratRange = caratRange
	}
	if price := c.Query("price"); price != "" {
		prices := strings.Split(price, "-")
		if len(prices) == 2 {
			minPrice, err := strconv.ParseFloat(prices[0], 64)
			if err == nil {
				filter.MinPrice = minPrice
			}
			maxPrice, err := strconv.ParseFloat(prices[1], 64)
			if err == nil {
				filter.MaxPrice = maxPrice
			}
		}
	}

	if pageSize := c.Query("page_size"); pageSize != "" {
		size, err := strconv.Atoi(pageSize)
		if err == nil {
			filter.PageSize = size
		}
	}

	if stoneCut := c.Query("stone_cut"); stoneCut != "" {
		filter.StoneCut = stoneCut
	}

	if page := c.Query("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err == nil {
			filter.Page = p
		}
	}

	products, err := cc.productService.GetProductsByCategoryID(int(categoryID), filter)
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	Render(c, prefix+"category.tmpl", gin.H{
		"CategoryID": categoryID,
		"Products":   products.Products,
		"Page":       products.Page,
		"Pages":      products.Pages,
		"Size":       products.Size,
		"Filter":     filter,
	})
}
