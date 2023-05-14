// routers/v1/user.go
package v1

import (
	v1 "swetelove/api/v1"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/categories", v1.GetCategories)
	group.GET("/category/:id/products", v1.GetProductsByCategory)
	group.GET("/advertisement/:position", v1.GetAdvertisementsByPosition)
	group.GET("/reviews/images", v1.GetProductReviewsWithImages)
	group.GET("/reviews/product/:product_id", v1.GetProductReviewsByProductID)
	group.POST("/register", v1.Register)
	group.POST("/login", v1.Login)
	group.GET("/product/:id", v1.GetProductByID)
	group.GET("/headerdata", v1.GetHeaderData)

	authenticated := group.Group("/")
	authenticated.Use(utils.AuthMiddleware())
	authenticated.GET("/user", v1.GetUser)
}
