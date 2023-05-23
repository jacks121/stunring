package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, template string, data interface{}) {
	client := c.Request.UserAgent()
	if strings.Contains(client, "Mobile") {
		template = "mobile/" + template
	} else {
		template = "pc/" + template
	}

	c.HTML(http.StatusOK, template, data)
}
