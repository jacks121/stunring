package routers

import (
	"fmt"
	"net/http"
	"swetelove/database"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func GlobalParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里设置全局参数
		commonService := service.NewCommonService(database.MysqlDB, database.RedisClient, database.GetContext())
		headerInfo, err := commonService.GetHeaderInfo()
		if err != nil {
			// Handle error here
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(headerInfo.CategoryTree) //在这里可以获取到数据
		c.Set("CategoryTree", headerInfo.CategoryTree)
		c.Set("Currency", headerInfo.Currencies)

		c.Next()
	}
}
