package collection

// CollectionType 定义产品集合类型
type CollectionType string

const (
	NewProducts CollectionType = "NewProducts"
	TopSales    CollectionType = "TopSales"
	Custom      CollectionType = "Custom"
	UserViewed  CollectionType = "UserViewed"
)
