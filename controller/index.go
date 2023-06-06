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
	// 从上下文中获取模板前缀
	prefix := c.MustGet("template_prefix").(string)

	// 获取顶部广告
	bannerAd, err := ic.AdService.GetAdvertisementByCode("banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取分类横幅广告
	indexCategoriesAd, err := ic.AdService.GetAdvertisementByCode("category_banner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取"新上市"产品集合
	newin, err := ic.CollectionService.GetProductsByCollectionCode("newin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取热销产品集合
	sales, err := ic.CollectionService.GetProductsByCollectionCode("top")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 渲染模板，并传递数据给模板
	Render(c, prefix+"index.tmpl", gin.H{
		"BannerAd":          bannerAd,
		"IndexCategoriesAd": indexCategoriesAd,
		"Newin":             newin,
		"Sales":             sales,
	})
}
