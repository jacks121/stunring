package routers

import (
	"net/http"
	"strconv"
	"swetelove/controller"
	"swetelove/repositories"
	"swetelove/utils"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Init constants
var (
	MobileUserAgent   string
	PCPrefix          string
	MobilePrefix      string
	IndexURL          string
	StaticURL         string
	StaticPath        string
	TemplatesPath     string
	TemplatePrefixKey string
)

func Init() {
	configPath := filepath.Join(".", "config.toml")
	config := utils.NewConfig(configPath)
	MobileUserAgent = config.Get("Template.MobileUserAgent")
	PCPrefix = config.Get("Template.PCPrefix")
	MobilePrefix = config.Get("Template.MobilePrefix")
	StaticURL = config.Get("Template.StaticURL")
	StaticPath = config.Get("Template.StaticPath")
	TemplatesPath = config.Get("Template.TemplatesPath")
	TemplatePrefixKey = config.Get("Template.TemplatePrefixKey")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.UseRawPath = true
	router.Use(ChangeTemplateBasedOnDevice())
	router.Delims("{!!", "!!}")
	router.LoadHTMLGlob(TemplatesPath)
	router.Static(StaticURL, StaticPath)

	indexController := controller.NewIndexController()
	cacheController := controller.NewCacheController()
	cacheProductController := controller.NewCacheProductController()
	ajaxController := controller.NewAjaxController()
	router.GET("/index", indexController.Index)
	router.GET("/", indexController.Index)
	router.GET("/cache/flush", cacheController.FlushCache)
	router.GET("/latest-products/:size", cacheProductController.GetLatestProducts)
	router.GET("/products/:id", cacheProductController.GetProductByID)
	router.GET("/review/gallery/", ajaxController.ReviewGallery)

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 获取所有的分类信息
		categoryRepo := repositories.NewCategoryRepository()
		categories, err := categoryRepo.GetAllCategories()
		if err != nil {
			// 处理错误
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 遍历分类信息，查找匹配的 URL
		var categoryID uint
		for _, category := range categories {
			if category.URL == path {
				categoryID = category.ID
				break
			}
		}

		categoryController := controller.NewCategoryController()
		// 如果找到匹配的分类 URL，则调用 categoryController 的 Show 方法并传递分类 ID
		if categoryID != 0 {
			c.Params = append(c.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(categoryID), 10)})
			categoryController.Show(c)
			return
		}

		// 处理其他逻辑，例如返回 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	return router
}

func ChangeTemplateBasedOnDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		if utils.IsMobileDevice(userAgent) {
			c.Set(TemplatePrefixKey, MobilePrefix)
		} else {
			c.Set(TemplatePrefixKey, PCPrefix)
		}

		c.Next()
	}
}
