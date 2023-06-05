package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	AdService         *service.AdvertisementService
	CollectionService *service.CollectionService
}

func NewIndexController() *IndexController {
	return &IndexController{
		AdService:         service.NewAdvertisementService(),
		CollectionService: service.NewCollectionService(),
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
	newin, err := ic.CollectionService.GetProductsByCollectionCode("newin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sales, err := ic.CollectionService.GetProductsByCollectionCode("top")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonNewin, err := json.MarshalIndent(sales, "", "  ")
	if err != nil {
		log.Println("Error marshaling newin to JSON:", err)
	} else {
		fmt.Println(string(jsonNewin))
	}

	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
		"Newin":             newin,
		"Sales":             sales,
	})
}
