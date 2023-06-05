package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	// 可以添加 CategoryService 或其他依赖的服务
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (cc *CategoryController) Show(c *gin.Context) {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// 根据 categoryID 执行相应的逻辑，例如查询分类信息，展示分类页面等
	// ...

	c.JSON(http.StatusOK, gin.H{
		"categoryID": categoryID,
	})
}
