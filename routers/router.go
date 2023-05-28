// routers/router.go
package routers

import (
	"strings"
	"swetelove/controller"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 将数据库连接添加到 Gin 上下文中
	router.Use(func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		// 根据 User-Agent 决定是移动设备还是PC
		if strings.Contains(userAgent, "Mobile") {
			c.Set("template_prefix", "mobile/")
		} else {
			c.Set("template_prefix", "pc/")
		}
	})

	router.Use(utils.CorsMiddleware())
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", controller.Index)
	router.Static("/static", "./static")

	return router
}
