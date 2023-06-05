package service

import (
	"swetelove/models"
	"swetelove/repositories"
	"swetelove/service/factory"
)

type CollectionService struct {
	CollectionRepository *repositories.CollectionRepository
}

func NewCollectionService() *CollectionService {
	return &CollectionService{
		CollectionRepository: repositories.NewCollectionRepository(),
	}
}

func (s *CollectionService) GetProductsByCollectionCode(code string) ([]models.Product, error) {
	collection, err := s.CollectionRepository.GetCollectionRule(code)
	if err != nil {
		return nil, err
	}

	factory := factory.NewStrategyFactory()
	strategy, err := factory.CreateStrategy(collection.Type, collection.Rule)
	if err != nil {
		return nil, err
	}

	products, err := strategy.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
