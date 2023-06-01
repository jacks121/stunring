package service

import (
	"swetelove/models"
	"swetelove/repositories"
)

type AdvertisementService struct {
	AdvertisementRepository *repositories.AdvertisementRepository // 广告仓库
}

// NewAdvertisementService 创建 AdvertisementService 实例
func NewAdvertisementService() *AdvertisementService {
	return &AdvertisementService{
		AdvertisementRepository: repositories.NewAdvertisementRepository(), // 初始化广告仓库
	}
}

// GetAdvertisementByCode 根据广告代码获取广告
func (s *AdvertisementService) GetAdvertisementByCode(code string) (*models.Advertisement, error) {
	return s.AdvertisementRepository.GetAdvertisementByCode(code)
}
