package repositories

import (
	"errors"
	"swetelove/models"
)

type SettingsRepository struct {
	BaseRepository
}

func NewSettingsRepository() *SettingsRepository {
	return &SettingsRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

// 获取所有Settings的方法
func (r *SettingsRepository) GetAllSettings() ([]models.Settings, error) {
	var settings []models.Settings
	redisKey := "settings:all"

	// 尝试从Redis获取数据
	if err := r.FetchFromRedis(redisKey, &settings); err != nil {
		// 如果从Redis获取数据失败，就从数据库中获取数据
		if err := r.FetchFromDB(&settings); err != nil {
			return nil, err
		}

		// 将获取到的数据写入Redis
		if err := r.StoreInRedis(redisKey, &settings); err != nil {
			return nil, err
		}
	}

	return settings, nil
}

// GetSettingByCode 方法，根据给定的代码获取特定的设置
func (r *SettingsRepository) GetSettingByCode(code string) (*models.Settings, error) {
	// 首先获取所有设置
	settings, err := r.GetAllSettings()
	if err != nil {
		return nil, err
	}

	// 遍历所有设置，查找特定的设置
	for _, setting := range settings {
		if setting.Code == code {
			return &setting, nil
		}
	}

	return nil, errors.New("setting not found")
}
