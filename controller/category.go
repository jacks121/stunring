package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	// Add CategoryService or other required services here
}

// NewCategoryController creates a new instance of CategoryController.
func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// Show handles the request to show a category.
func (cc *CategoryController) Show(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	Render(c, prefix+"category.tmpl", gin.H{
		"CategoryID": categoryID,
	})
}
