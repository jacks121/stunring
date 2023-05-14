package v1

import (
	"net/http"
	"strconv"
	"swetelove/services"
	"swetelove/utils"

	"github.com/gin-gonic/gin"
)

func GetProductReviewsByProductID(c *gin.Context) {
	// 从路由参数中获取商品ID
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		response := utils.NewApiResponse("error", "Invalid product ID", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// 从上下文中获取数据库和Redis客户端
	db, redisClient, err := getDatabaseAndRedis(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建ProductReviewService服务
	productReviewService, err := services.NewProductReviewService(db, redisClient)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to create product review service", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	// 调用服务获取商品ID对应的评论列表
	reviews, err := productReviewService.GetReviewsByProductID(int(productID), page, perPage)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch product reviews", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 构造响应并返回
	response := utils.NewApiResponse("success", "Product reviews fetched successfully", reviews)
	c.JSON(http.StatusOK, response)
}

func GetProductReviewsWithImages(c *gin.Context) {
	// 从上下文中获取数据库和Redis客户端
	db, redisClient, err := getDatabaseAndRedis(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建ProductReviewService服务
	productReviewService, err := services.NewProductReviewService(db, redisClient)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to create product review service", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	// 调用服务获取带图片的评论列表
	reviews, err := productReviewService.GetReviewsWithImages(page, perPage)
	if err != nil {
		response := utils.NewApiResponse("error", "Failed to fetch product reviews with images", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 构造响应并返回
	response := utils.NewApiResponse("success", "Product reviews with images fetched successfully", reviews)
	c.JSON(http.StatusOK, response)
}
