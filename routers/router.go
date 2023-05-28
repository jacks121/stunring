package routers

import (
	"fmt"
	"net/http"
	"swetelove/controller"
	"swetelove/database"
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

	router.LoadHTMLGlob(TemplatesPath)
	router.Static(StaticURL, StaticPath)

	router.GET("/index", controller.Index)
	router.GET("/", controller.Index)

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 获取所有的分类信息
		categoryRepo := &repositories.CategoryRepository{
			DB:    database.MysqlDB,
			Redis: database.RedisClient,
			Ctx:   database.GetContext(),
		}
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
		fmt.Println("xxxxxxxxxxxx")
		fmt.Println(categoryID)
		// 如果找到匹配的分类 URL，则调用 categoryShow 处理函数并传递分类 ID
		if categoryID != 0 {
			controller.CategoryShow(c, categoryID)
			return
		}
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
