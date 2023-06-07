package controller

import (
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
	// Get the template prefix from the context
	prefix := c.MustGet("template_prefix").(string)

	// Get the top banner advertisement
	bannerAd, err := ic.AdService.GetAdvertisementByCode("banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the category banner advertisement
	indexCategoriesAd, err := ic.AdService.GetAdvertisementByCode("category_banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the "New In" product collection
	newin, err := ic.CollectionService.GetProductsByCollectionCode("newin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the top-selling product collection
	sales, err := ic.CollectionService.GetProductsByCollectionCode("top")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template and pass data to the template
	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
		"Newin":             newin,
		"Sales":             sales,
	})
}
