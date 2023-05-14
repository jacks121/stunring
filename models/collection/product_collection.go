package collection

import "swetelove/models"

// ProductCollection 接口定义获取产品列表的通用方法
type ProductCollection interface {
	GetCollection() ([]models.Product, error)
}
