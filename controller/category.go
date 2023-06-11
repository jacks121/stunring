package controller

import (
	"net/http"
	"strconv"
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

	products, err := cc.productService.GetProductsByCategoryID(int(categoryID))
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	Render(c, prefix+"category.tmpl", gin.H{
		"CategoryID": categoryID,
		"Products":   products,
	})
}
