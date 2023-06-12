package controller

import (
	"net/http"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

// Render renders the HTML template with the given data and adds header information.
func Render(c *gin.Context, templateName string, data gin.H) {
	commonService := service.NewCommonService()
	headerInfo, err := commonService.GetHeaderInfo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data["HeaderInfo"] = headerInfo
	data["BaseUrl"] = "//192.168.183.209:8080"
	c.HTML(http.StatusOK, templateName, data)
}
