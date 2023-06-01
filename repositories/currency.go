package repositories

import (
	"swetelove/models"
)

type CurrencyRepository struct {
	BaseRepository
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

func (r *CurrencyRepository) GetAllCurrencies() ([]*models.Currency, error) {
	var currencies []*models.Currency
	redisKey := "currencies"
	if err := r.FetchFromRedis(redisKey, &currencies); err != nil {
		// 如果从Redis获取数据失败，就从数据库中获取数据
		if err := r.FetchFromDB(&currencies); err != nil {
			return nil, err
		}
		// 将获取到的数据写入Redis
		if err := r.StoreInRedis(redisKey, &currencies); err != nil {
			return nil, err
		}
	}
	return currencies, nil
}
