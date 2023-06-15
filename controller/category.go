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
	BaseController
	productService     *service.ProductService
	settingsRepository *repositories.SettingsRepository
}

// 创建 CategoryController 实例
func NewCategoryController() *CategoryController {
	return &CategoryController{
		productService:     service.NewProductService(),
		settingsRepository: repositories.NewSettingsRepository(),
	}
}

// 将查询字符串转化为int类型的函数
func queryStringToInt(query string) int {
	num, _ := strconv.Atoi(query)
	return num
}

// 将价格区间转化为数组的函数
func parsePriceRange(price string) [2]int {
	priceRangeStr := strings.Split(price, "-")
	priceRange := [2]int{}
	if len(priceRangeStr) == 2 {
		priceRange[0] = queryStringToInt(priceRangeStr[0])
		priceRange[1] = queryStringToInt(priceRangeStr[1])
	}
	return priceRange
}

func (cc *CategoryController) Show(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)

	// 获取分类ID，并进行错误处理
	categoryID := queryStringToInt(c.Param("id"))

	// 处理分页的大小，如果没有指定则使用默认值
	pageSize := queryStringToInt(c.Query("page_size"))
	if pageSize == 0 {
		pageSize = 2 // 默认每页显示24个商品
	}

	// 创建ProductListParams
	params := repositories.ProductListParams{
		SortBy:     c.Query("product_list_order"),
		Page:       queryStringToInt(c.Query("p")),
		PageSize:   pageSize,
		PriceRange: parsePriceRange(c.Query("price")),
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

	// 获取 code 为 filter 的 settings
	settings, err := cc.settingsRepository.GetSettingByCode("filter")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 创建一个 map 来保存查询参数
	getParams := make(map[string]string)

	// 获取所有查询参数
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		// 只保存每个参数的第一个值
		if len(values) > 0 {
			getParams[key] = values[0]
		}
	}
	fmt.Println(settings)
	Render(c, prefix+"category.tmpl", gin.H{
		"CategoryID":  categoryID,
		"Products":    result.Products,
		"Params":      params,
		"Size":        result.Size,
		"CurrentPage": result.CurrentPage,
		"TotalPages":  result.TotalPages,
		"Url":         c.Request.Host + c.Request.URL.Path,
		"RawQuery":    c.Request.URL.RawQuery,
		"GetParams":   getParams,
		"Settings":    settings,
	})
}
