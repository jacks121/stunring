package repositories

import (
	"fmt"
	"swetelove/models"
)

type AdvertisementRepository struct {
	BaseRepository
}

func NewAdvertisementRepository() *AdvertisementRepository {
	return &AdvertisementRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

func (r *AdvertisementRepository) GetAdvertisementByCode(code string) (*models.Advertisement, error) {
	var ad models.Advertisement
	redisKey := fmt.Sprintf("advertisement:%v", code)
	if err := r.FetchFromRedis(redisKey, &ad); err != nil {
		// 如果从Redis获取数据失败，就从数据库中获取数据
		if err := r.FetchFromDB(&ad, code); err != nil {
			return nil, err
		}
		// 将获取到的数据写入Redis
		if err := r.StoreInRedis(redisKey, &ad); err != nil {
			return nil, err
		}
	}
	return &ad, nil
}

func (r *AdvertisementRepository) FetchFromDB(ad *models.Advertisement, code string) error {
	// 在这里，我们可以添加特定的查询逻辑，比如预加载Images
	return r.DB.Preload("Images").Where("code = ?", code).First(ad).Error
}
