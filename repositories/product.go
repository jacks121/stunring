package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"swetelove/database"
	"swetelove/models"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	BaseRepository
	ES *elastic.Client
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		BaseRepository: *NewBaseRepository(),
		ES:             database.ElasticClient,
	}
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := r.DB.Preload(clause.Associations).
		Preload("ProductAttributes.Value").
		Preload("ProductAttributes.Images").
		Preload("Reviews", func(db *gorm.DB) *gorm.DB {
			return db.Order("rating DESC").Limit(5)
		}).
		Preload("Reviews.Images").
		Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to get products from DB: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product

	// if true {
	// 	if err := r.DB.Preload(clause.Associations).Preload("Reviews", func(db *gorm.DB) *gorm.DB {
	// 		return db.Order("rating DESC").Limit(5)
	// 	}).Preload("Reviews.Images").First(&product, id).Error; err != nil {
	// 		return nil, fmt.Errorf("failed to get product from DB: %w", err)
	// 	}
	// 	return &product, nil
	// }

	// 1. Try to fetch from Redis
	redisKey := fmt.Sprintf("product:%d", id)
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product from Redis: %w", err)
		}
		return &product, nil
	}

	// 2. Try to fetch from Elasticsearch
	esRes, err := r.ES.Get().Index("products").Id(strconv.Itoa(int(id))).Do(r.Ctx)
	if err == nil && esRes.Found {
		if err = json.Unmarshal(esRes.Source, &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product from Elasticsearch: %w", err)
		}
		// Update Redis with a random expiration time between 600 to 1200 minutes
		expiration := time.Duration(rand.Intn(601)+600) * time.Minute
		redisVal, err := json.Marshal(product)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal product for Redis: %w", err)
		}
		if err = r.Redis.Set(r.Ctx, redisKey, redisVal, expiration).Err(); err != nil {
			return nil, fmt.Errorf("failed to set product in Redis: %w", err)
		}
		return &product, nil
	}

	// 3. Fetch from MySQL
	if err = r.DB.Preload(clause.Associations).Preload("ProductAttributes.Value").
		Preload("ProductAttributes.Images").Preload("Reviews", func(db *gorm.DB) *gorm.DB {
		return db.Order("rating DESC").Limit(5)
	}).Preload("Reviews.Images").First(&product, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get product from DB: %w", err)
	}

	// Update Redis and Elasticsearch
	expiration := time.Duration(rand.Intn(601)+600) * time.Minute
	redisVal, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, expiration).Err(); err != nil {
		return nil, fmt.Errorf("failed to set product in Redis: %w", err)
	}
	_, err = r.ES.Index().Index("products").Id(strconv.Itoa(int(id))).BodyJson(product).Do(r.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set product in Elasticsearch: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) GetLatestProducts(size int, sortBy string) ([]models.Product, error) {
	// 创建 Elasticsearch 搜索请求
	return r.GetLatestProductsFromDB(size, sortBy)
	searchRequest := r.ES.Search().Index("products")

	// 根据 sortBy 参数设置排序字段和顺序
	switch sortBy {
	case "created_at":
		searchRequest.Sort("CreatedAt", false) // 根据 created_at 字段降序排列
	case "sales":
		searchRequest.Sort("Sales", false) // 根据 sales 字段降序排列
	default:
		return nil, fmt.Errorf("invalid sortBy parameter: %s", sortBy)
	}

	// 设置查询结果数量
	searchRequest.Size(size)

	// 发送搜索请求
	searchResult, err := searchRequest.Do(database.GetContext())
	if err != nil || len(searchResult.Hits.Hits) == 0 {
		// If Elasticsearch request fails or there are no results, fallback to DB
		return r.GetLatestProductsFromDB(size, sortBy)
	}

	// 解析搜索结果
	var products []models.Product
	for _, hit := range searchResult.Hits.Hits {
		var product models.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			// If unmarshalling fails, fallback to DB
			return r.GetLatestProductsFromDB(size, sortBy)
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetLatestProductsFromDB(size int, sortBy string) ([]models.Product, error) {
	var products []models.Product

	var order string
	switch sortBy {
	case "created_at":
		order = "created_at DESC" // 根据 created_at 字段降序排列
	case "sales":
		order = "sales DESC" // 根据 sales 字段降序排列
	default:
		return nil, fmt.Errorf("invalid sortBy parameter: %s", sortBy)
	}

	if err := r.DB.Preload(clause.Associations).
		Preload("ProductAttributes.Value").
		Preload("ProductAttributes.Images").
		Preload("Reviews", func(db *gorm.DB) *gorm.DB {
			return db.Order("rating DESC").Limit(5)
		}).
		Preload("Reviews.Images").
		Order(order).Limit(size).
		Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to get products from DB: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) SyncProductsToES() error {
	// 删除目标索引
	_, err := r.ES.DeleteIndex("products").Do(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting index: %v", err)
	}

	// 从数据库获取所有商品
	products, err := r.GetAllProducts()
	if err != nil {
		return fmt.Errorf("error getting products from DB: %v", err)
	}

	var wg sync.WaitGroup
	for _, product := range products {
		wg.Add(1)
		go func(p models.Product) {
			defer wg.Done()

			_, err := r.ES.Index().
				Index("products").
				Id(fmt.Sprint(p.ID)).
				BodyJson(p).
				Do(context.Background())
			if err != nil {
				// 错误处理，例如记录日志或者将错误发送到错误追踪系统
				fmt.Printf("error indexing product %v: %v", p.ID, err)
			}
		}(product)
	}
	wg.Wait()

	return nil
}
