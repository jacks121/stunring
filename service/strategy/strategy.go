package strategy

import (
	"swetelove/models"
	"swetelove/repositories"
)

type Strategy interface {
	GetProducts() ([]models.Product, error)
}

type NewInStrategy struct {
	ProductRepository *repositories.ProductRepository
	Limit             int
}

func (s *NewInStrategy) GetProducts() ([]models.Product, error) {
	products, err := s.ProductRepository.GetLatestProducts(s.Limit, "created_at")
	if err != nil {
		return nil, err
	}
	return products, nil
}

type TopSellersStrategy struct {
	ProductRepository *repositories.ProductRepository
	Limit             int
}

func (s *TopSellersStrategy) GetProducts() ([]models.Product, error) {
	products, err := s.ProductRepository.GetLatestProducts(s.Limit, "sales")
	if err != nil {
		return nil, err
	}
	return products, nil
}
