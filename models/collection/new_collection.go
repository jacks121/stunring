package collection

import (
	"swetelove/models"
	"swetelove/repositories"
)

type NewCollection struct {
	Count       int
	productRepo *repositories.ProductRepository
}

func (c *NewCollection) GetCollection() ([]models.Product, error) {
	collection, _ := c.productRepo.GetNewProducts(c.Count)
	return collection, nil
}
