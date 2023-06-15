package controller

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (bc *BaseController) jsonError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
