// routers/router.go
package routers

import (
	v1 "swetelove/routers/v1"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, redisClient *redis.Client, config *utils.Config) *gin.Engine {
	router := gin.Default()

	// 将数据库连接添加到 Gin 上下文中
	router.Use(func(c *gin.Context) {
		c.Set("redis", redisClient)
		c.Set("db", db)
		c.Set("config", config)
		c.Next()
	})

	router.Use(utils.CorsMiddleware())

	apiV1 := router.Group("/api/v1")

	v1.RegisterRoutes(apiV1)

	return router
}
