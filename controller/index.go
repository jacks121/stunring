package controller

import (
	"net/http"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	AdService *service.AdvertisementService
}

func NewIndexController() *IndexController {
	return &IndexController{
		AdService: service.NewAdvertisementService(),
	}
}

func (ic *IndexController) Index(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	bannerAd, err := ic.AdService.GetAdvertisementByCode("banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	indexCategoriesAd, err := ic.AdService.GetAdvertisementByCode("category_banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
	})
}
