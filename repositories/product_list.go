package repositories

import (
	"context"
	"encoding/json"
	"swetelove/models"

	"github.com/olivere/elastic/v7"
)

type ProductListResult struct {
	Products    []models.Product `json:"products"`
	TotalPages  int
	CurrentPage int
	Size        int
}

type ProductListParams struct {
	SortBy     string
	SortOrder  string
	PriceRange [2]int
	Filters    map[string]string
	Page       int
	PageSize   int
}

func (r *ProductRepository) GetProductsByCategoryID(categoryID int, params ProductListParams) (ProductListResult, error) {
	ctx := context.Background()

	// 构建查询条件，匹配指定的Categories ID
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("Categories.ID", categoryID))

	// 添加价格筛选条件
	if params.PriceRange[0] > 0 || params.PriceRange[1] > 0 {
		priceRangeQuery := elastic.NewRangeQuery("CurrentPrice").
			Gte(params.PriceRange[0]).
			Lte(params.PriceRange[1])
		boolQuery.Filter(priceRangeQuery)
	}

	// 添加其他筛选条件
	for field, value := range params.Filters {
		if value != "" {
			boolQuery.Filter(elastic.NewTermQuery(field, value))
		}
	}

	if params.Page <= 0 {
		params.Page = 1 // 设置默认的页码为1
	}

	if params.PageSize <= 0 {
		params.PageSize = 24 // 设置默认的页面大小为24
	}

	// 计算从哪个位置开始返回结果
	from := (params.Page - 1) * params.PageSize
	// 构建搜索请求
	searchService := r.ES.Search().
		Index("products").    // 设置索引名称为 "products"
		Query(boolQuery).     // 设置查询条件
		From(from).           // 设置从哪个位置开始返回结果
		Size(params.PageSize) // 设置返回的最大文档数量

	// 设置排序条件
	if params.SortBy != "" && params.SortOrder != "" {
		searchService.Sort(params.SortBy, params.SortOrder == "asc")
	}

	searchResult, err := searchService.Do(ctx) // 执行搜索

	if err != nil {
		// 处理搜索错误
		return ProductListResult{}, err
	}

	// 计算总页数
	total := int(searchResult.Hits.TotalHits.Value)
	pages := total / params.PageSize
	if total%params.PageSize > 0 {
		pages++
	}

	// 解析搜索结果
	var products []models.Product
	for _, hit := range searchResult.Hits.Hits {
		var product models.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			// 处理解析错误
			return ProductListResult{}, err
		}
		products = append(products, product)
	}

	return ProductListResult{
		Products:    products,
		TotalPages:  pages,
		CurrentPage: params.Page,
		Size:        params.PageSize,
	}, nil
}
