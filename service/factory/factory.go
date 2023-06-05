package factory

import (
	"encoding/json"
	"fmt"

	"swetelove/repositories"
	"swetelove/service/strategy"
)

type StrategyFactory struct {
	ProductRepository *repositories.ProductRepository
}

func NewStrategyFactory() *StrategyFactory {
	return &StrategyFactory{
		ProductRepository: repositories.NewProductRepository(),
	}
}

func (f *StrategyFactory) CreateStrategy(collectionType string, rule json.RawMessage) (strategy.Strategy, error) {
	switch collectionType {
	case "new":
		var params struct {
			Limit int `json:"limit"`
		}
		json.Unmarshal(rule, &params)
		return &strategy.NewInStrategy{
			ProductRepository: f.ProductRepository,
			Limit:             params.Limit,
		}, nil
	case "sales":
		var params struct {
			Limit int `json:"limit"`
		}
		json.Unmarshal(rule, &params)
		return &strategy.TopSellersStrategy{
			ProductRepository: f.ProductRepository,
			Limit:             params.Limit,
		}, nil
	default:
		return nil, fmt.Errorf("unknown strategy code: %s", collectionType)
	}
}
