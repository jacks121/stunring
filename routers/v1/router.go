// routers/v1/user.go
package v1

import (
	"context"
	"net/http"
	"strconv"
	"swetelove/database"
	"swetelove/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/product/:id", func(c *gin.Context) {
		db := database.MysqlDB
		redisClient := database.RedisClient
		elasticClient := database.ElasticClient
		ctx := context.Background()

		repo := repositories.NewProductRepository(db, redisClient, elasticClient, ctx)

		// Get the ID from the URL
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		// Call the GetProductByID method
		product, err := repo.GetProductByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// If there is no error, return the product
		c.JSON(http.StatusOK, product)
	})
	// group.GET("/category/:id/products", v1.GetProductsByCategory)
	// group.GET("/advertisement/:position", v1.GetAdvertisementsByPosition)
	// group.GET("/reviews/images", v1.GetProductReviewsWithImages)
	// group.GET("/reviews/product/:product_id", v1.GetProductReviewsByProductID)
	// group.POST("/register", v1.Register)
	// group.POST("/login", v1.Login)
	// group.GET("/product/:id", v1.GetProductByID)
	// group.GET("/headerdata", v1.GetHeaderData)
	// group.GET("/products/index", v1.IndexProductsToElasticsearch)

	// authenticated := group.Group("/")
	// authenticated.Use(utils.AuthMiddleware())
	// authenticated.GET("/user", v1.GetUser)
}
