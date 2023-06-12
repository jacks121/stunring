package controller

import (
	"fmt"
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

func (cc *CategoryController) Show(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	// 获取分类ID
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid category ID"})
		return
	}

	// 创建ProductListParams
	page, _ := strconv.Atoi(c.Query("p"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize == 0 {
		pageSize = 2 // 默认每页显示24个商品
	}
	priceRangeStr := strings.Split(c.Query("price"), "-")
	priceRange := [2]int{}
	if len(priceRangeStr) == 2 {
		priceRange[0], _ = strconv.Atoi(priceRangeStr[0])
		priceRange[1], _ = strconv.Atoi(priceRangeStr[1])
	}
	params := repositories.ProductListParams{
		SortBy:     c.Query("product_list_order"),
		Page:       page,
		PageSize:   pageSize,
		PriceRange: priceRange,
		Filters: map[string]string{
			"stone_cut":     c.Query("stone_cut"),
			"stone_color":   c.Query("stone_color"),
			"ring_crafting": c.Query("ring_crafting"),
		},
	}

	// 调用GetProductsByCategoryID函数
	result, err := cc.productService.GetProductsByCategoryID(categoryID, params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	startPage := params.Page - 2
	if startPage < 1 {
		startPage = 1
	}
	endPage := startPage + 4
	if endPage > result.TotalPages {
		endPage = result.TotalPages
	}
	pages := make([]int, endPage-startPage+1)
	for i := range pages {
		pages[i] = startPage + i
	}
	nextPage := params.Page + 1
	if nextPage > result.TotalPages {
		nextPage = params.Page
	}
	fmt.Println("CurrentPage:", result.CurrentPage)
	fmt.Println("TotalPages:", result.TotalPages)
	fmt.Println("Pages:", pages)
	fmt.Println("NextPage:", nextPage)

	Render(c, prefix+"category.tmpl", gin.H{
		"CategoryID":  categoryID,
		"Products":    result.Products,
		"Params":      params,
		"Size":        result.Size,
		"CurrentPage": result.CurrentPage,
		"TotalPages":  result.TotalPages,
		"Pages":       pages,
		"NextPage":    nextPage,
		"Url":         c.Request.Host + c.Request.URL.Path,
		"RawQuery":    c.Request.URL.RawQuery,
	})
}
