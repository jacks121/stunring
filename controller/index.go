// controller/index.go
package controller

import (
	"net/http"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	adService := service.NewAdvertisementService()
	bannerAd, err := adService.GetAdvertisementByCode("banner")
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	indexCategoriesAd, err := adService.GetAdvertisementByCode("category_banner")
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
	})
}
