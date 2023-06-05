package service

import (
	"fmt"
	"swetelove/repositories"
)

type CacheService struct {
	productRepository *repositories.ProductRepository
}

func NewCacheService() *CacheService {
	return &CacheService{
		productRepository: repositories.NewProductRepository(),
	}
}

func (s *CacheService) SyncProductsToES() error {
	err := s.productRepository.SyncProductsToES()
	if err != nil {
		return fmt.Errorf("error syncing products to ES: %v", err)
	}
	return nil
}
