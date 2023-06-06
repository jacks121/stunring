package routers

import (
	"net/http"
	"strconv"
	"swetelove/controller"
	"swetelove/repositories"
	"swetelove/utils"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Init constants
var (
	MobileUserAgent   string
	PCPrefix          string
	MobilePrefix      string
	IndexURL          string
	StaticURL         string
	StaticPath        string
	TemplatesPath     string
	TemplatePrefixKey string
)

func Init() {
	configPath := filepath.Join(".", "config.toml")
	config := utils.NewConfig(configPath)
	MobileUserAgent = config.Get("Template.MobileUserAgent")
	PCPrefix = config.Get("Template.PCPrefix")
	MobilePrefix = config.Get("Template.MobilePrefix")
	StaticURL = config.Get("Template.StaticURL")
	StaticPath = config.Get("Template.StaticPath")
	TemplatesPath = config.Get("Template.TemplatesPath")
	TemplatePrefixKey = config.Get("Template.TemplatePrefixKey")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.UseRawPath = true
	router.Use(ChangeTemplateBasedOnDevice())
	router.Delims("{!!", "!!}")
	router.LoadHTMLGlob(TemplatesPath)
	router.Static(StaticURL, StaticPath)

	indexController := controller.NewIndexController()
	cacheController := controller.NewCacheController()
	cacheProductController := controller.NewCacheProductController()
	router.GET("/index", indexController.Index)
	router.GET("/", indexController.Index)
	router.GET("/cache/flush", cacheController.FlushCache)
	router.GET("/latest-products/:size", cacheProductController.GetLatestProducts)
	router.GET("/products/:id", cacheProductController.GetProductByID)
	router.GET("/review/gallery/", func(c *gin.Context) {
		data := gin.H{
			"items": []gin.H{
				{
					"title":      "shiny and gorgeous",
					"src":        "https://cdn.stunring.com/media/review_images/d/i/dingtalk_20230602144445.jpg",
					"productUrl": "https://www.stunring.com/unique-29ct-sapphire-stones-radiant-cut-engagement-ring-2130410.html",
				},
				{
					"title":      "LOVE",
					"src":        "https://cdn.stunring.com/media/review_images/2/5/2530537-5.jpg",
					"productUrl": "https://www.stunring.com/stunring-gorgeous-animal-snake-shape-emerald-green-eyes-pear-cut-ring-2530537.html",
				},
				{
					"title":      "Great",
					"src":        "https://cdn.stunring.com/media/review_images/_/_/__20230427154032.png",
					"productUrl": "https://www.stunring.com/stunring-gorgeous-snake-shape-bracelet-21130429.html",
				},
				{
					"title":      "Good",
					"src":        "https://cdn.stunring.com/media/review_images/2/5/2530537-6.jpg",
					"productUrl": "https://www.stunring.com/stunring-gorgeous-animal-snake-shape-emerald-green-eyes-pear-cut-ring-2530537.html",
				},
				{
					"title":      "Boom!",
					"src":        "https://cdn.stunring.com/media/review_images/_/1/_1_.jpg",
					"productUrl": "https://www.stunring.com/stunring-handmade-summer-snake-lab-created-white-gemstone-fine-necklace-nk015.html",
				},
				{
					"title":      "stunning",
					"src":        "https://cdn.stunring.com/media/review_images/l/q/lqlpjw3dovt4wsjna8dnahywbk6ltad-pcaeeidinodmaa_540_960.png",
					"productUrl": "https://www.stunring.com/stunring-2-0ct-independent-design-oval-cut-ruby-engagement-ring-ek051.html",
				},
				{
					"title":      "gorgeous!",
					"src":        "https://cdn.stunring.com/media/review_images/2/8/284725477_10223201818608621_2801487919247514155_n.jpg",
					"productUrl": "https://www.stunring.com/stunring-1-13-ct-art-deco-gem-and-round-vintage-engagement-ring-set-ek045.html",
				},
				{
					"title":      "Good!",
					"src":        "https://cdn.stunring.com/media/review_images/2/5/2530537-8.jpg",
					"productUrl": "https://www.stunring.com/stunring-gorgeous-animal-snake-shape-emerald-green-eyes-pear-cut-ring-2530537.html",
				},
				{
					"title":      "so beautiful",
					"src":        "https://cdn.stunring.com/media/review_images/_/1/_1080_1920.png",
					"productUrl": "https://www.stunring.com/stunring-oval-cut-eternity-cornflower-aquamarine-gem-and-fashion-sweetie-band-ring-bk035.html",
				},
				{
					"title":      "Beautiful",
					"src":        "https://cdn.stunring.com/media/review_images/1/_/1.jpg",
					"productUrl": "https://www.stunring.com/stunring-handmade-summer-snake-lab-created-white-gemstone-fine-necklace-nk015.html",
				},
			},
			"hasMorePages": true,
		}

		c.JSON(http.StatusOK, data)
	})

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 获取所有的分类信息
		categoryRepo := repositories.NewCategoryRepository()
		categories, err := categoryRepo.GetAllCategories()
		if err != nil {
			// 处理错误
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 遍历分类信息，查找匹配的 URL
		var categoryID uint
		for _, category := range categories {
			if category.URL == path {
				categoryID = category.ID
				break
			}
		}

		categoryController := controller.NewCategoryController()
		// 如果找到匹配的分类 URL，则调用 categoryController 的 Show 方法并传递分类 ID
		if categoryID != 0 {
			c.Params = append(c.Params, gin.Param{Key: "id", Value: strconv.FormatUint(uint64(categoryID), 10)})
			categoryController.Show(c)
			return
		}

		// 处理其他逻辑，例如返回 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	return router
}

func ChangeTemplateBasedOnDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		if utils.IsMobileDevice(userAgent) {
			c.Set(TemplatePrefixKey, MobilePrefix)
		} else {
			c.Set(TemplatePrefixKey, PCPrefix)
		}

		c.Next()
	}
}
