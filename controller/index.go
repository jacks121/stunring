// controller/index.go
package controller

import (
	"fmt"
	"net/http"
	"swetelove/database"
	"swetelove/service"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	prefix := c.MustGet("template_prefix").(string)
	adService := service.NewAdvertisementService(database.MysqlDB, database.RedisClient, database.GetContext())
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
	fmt.Printf("Person: %+v\n", bannerAd)
	fmt.Printf("Person: %+v\n", indexCategoriesAd)
	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
	})
}
