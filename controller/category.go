// controller/index.go
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryShow(c *gin.Context, categoryID uint) {
	// 根据 categoryID 执行相应的逻辑，例如查询分类信息，展示分类页面等
	// ...

	c.JSON(http.StatusOK, gin.H{
		"categoryID": categoryID,
	})
}
