// controller/index.go
package controller

import (
	"net/http"
	"swetelove/database"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	commonService := service.NewCommonService(database.MysqlDB, database.RedisClient, database.GetContext())
	headerInfo, err := commonService.GetHeaderInfo()

	if err != nil {
		// Handle error here
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, prefix+"index.tmpl", gin.H{
		"CategoryTree": headerInfo.CategoryTree,
		"Currency":     headerInfo.Currencies,
	})
}
