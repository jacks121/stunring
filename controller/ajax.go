package controller

import (
	"net/http"
	"strconv"
	"swetelove/repositories"

	"github.com/gin-gonic/gin"
)

type AjaxController struct {
	reviewRepo *repositories.ReviewRepository
}

func NewAjaxController() *AjaxController {
	return &AjaxController{
		reviewRepo: repositories.NewReviewRepository(),
	}
}

type ReviewItem struct {
	Title      string `json:"title"`
	Src        string `json:"src"`
	ProductURL string `json:"productUrl"`
}

type ReviewResponse struct {
	Items        []ReviewItem `json:"items"`
	HasMorePages bool         `json:"hasMorePages"`
}

// ReviewGallery handles the review gallery AJAX request.
func (ac *AjaxController) ReviewGallery(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("nextPage")) // Get the page number from the request parameter
	if err != nil {
		page = 1 // Default to the first page
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10")) // Get the page size from the request parameter, default to 10
	if err != nil {
		pageSize = 10 // Default page size is 10
	}

	// Call the GetReviewsWithImagesByPageAndCount method of ReviewRepository to retrieve review data
	reviews, total, err := ac.reviewRepo.GetReviewsWithImagesByPageAndCount(page, pageSize)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct the response data structure
	response := ReviewResponse{
		Items:        make([]ReviewItem, len(reviews)),
		HasMorePages: (page * pageSize) < total, // Check if there are more pages
	}

	// Convert each review data to the desired format
	for i, review := range reviews {
		imageURL := ""
		if len(review.Images) > 0 {
			imageURL = review.Images[0].ImageURL // Get the URL of the first image
		}

		response.Items[i] = ReviewItem{
			Title:      "",
			Src:        imageURL,
			ProductURL: "",
		}
	}

	// Return the JSON data
	c.JSON(http.StatusOK, response)
}
