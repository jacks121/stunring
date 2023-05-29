// controller/index.go
package controller

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)

	Render(c, prefix+"index.tmpl", gin.H{})
}
