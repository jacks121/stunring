package routers

import (
	"swetelove/controller"
	"swetelove/utils" // 假设这是一个新的工具包，其中包含UserAgent检测功能

	"github.com/gin-gonic/gin"
)

// 设定常量
const (
	MobileUserAgent   = "Mobile"
	PCPrefix          = "pc/"
	MobilePrefix      = "mobile/"
	IndexURL          = "/index"
	StaticURL         = "/static"
	StaticPath        = "./static"
	TemplatesPath     = "templates/**/*"
	TemplatePrefixKey = "template_prefix"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(ChangeTemplateBasedOnDevice())

	router.LoadHTMLGlob(TemplatesPath)
	router.Static(StaticURL, StaticPath)

	router.GET(IndexURL, controller.RenderIndexPage)

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
