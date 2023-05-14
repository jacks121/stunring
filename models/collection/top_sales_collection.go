package collection

import (
	"swetelove/models"
	"swetelove/repositories"
)

type TopSalesCollection struct {
	Count       int
	productRepo *repositories.ProductRepository
}

func (c *TopSalesCollection) GetCollection() ([]models.Product, error) {
	collection, _ := c.productRepo.GetNewProducts(c.Count)
	return collection, nil
}
