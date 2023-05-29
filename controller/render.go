package controller

import (
	"net/http"
	"swetelove/database"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, templateName string, data gin.H) {
	commonService := service.NewCommonService(database.MysqlDB, database.RedisClient, database.GetContext())
	headerInfo, err := commonService.GetHeaderInfo()

	if err != nil {
		// Handle error here
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data["HeaderInfo"] = headerInfo
	c.HTML(http.StatusOK, templateName, data)
}
