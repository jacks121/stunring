package service

import (
	"swetelove/models"
	"swetelove/repositories"
)

// HeaderInfo 包含头部信息
type HeaderInfo struct {
	CategoryTree []*models.Category // 类别树
	Currencies   []*models.Currency // 货币列表
	// 其他字段...
}

// CommonService 提供通用服务
type CommonService struct {
	CategoryRepository *repositories.CategoryRepository // 类别仓库
	CurrencyRepository *repositories.CurrencyRepository // 货币仓库
	// 其他仓库...
}

// NewCommonService 创建CommonService实例
func NewCommonService() *CommonService {
	return &CommonService{
		CategoryRepository: repositories.NewCategoryRepository(), // 初始化类别仓库
		CurrencyRepository: repositories.NewCurrencyRepository(), // 初始化货币仓库
		// 初始化其他仓库...
	}
}

// GetHeaderInfo 获取头部信息
func (s *CommonService) GetHeaderInfo() (*HeaderInfo, error) {
	info := &HeaderInfo{}

	// 获取类别树
	categoryTree, err := s.CategoryRepository.GetCategoryTree()
	if err != nil {
		return nil, err
	}
	info.CategoryTree = categoryTree

	// 获取货币列表
	currencies, err := s.CurrencyRepository.GetAllCurrencies()
	if err != nil {
		return nil, err
	}
	info.Currencies = currencies

	// 获取其他信息...
	// ...

	return info, nil
}
