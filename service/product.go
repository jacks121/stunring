package service

import (
	"swetelove/models"
	"swetelove/repositories"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository // 商品仓库
}

// NewProductService 创建 ProductService 实例
func NewProductService() *ProductService {
	return &ProductService{
		ProductRepository: repositories.NewProductRepository(), // 初始化商品仓库
	}
}

// GetProductsByCategoryID 根据分类ID获取商品列表
func (s *ProductService) GetProductsByCategoryID(categoryID int) ([]models.Product, error) {
	return s.ProductRepository.GetProductsByCategoryID(categoryID)
}
