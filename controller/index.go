// controller/index.go
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderIndexPage(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	c.HTML(http.StatusOK, prefix+"index.tmpl", gin.H{
		"title": "首页",
	})
}
