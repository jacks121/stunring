package models

type ProductCategory struct {
	ProductID  uint `gorm:"primary_key;auto_increment:false" json:"product_id"`
	CategoryID uint `gorm:"primary_key;auto_increment:false" json:"category_id"`
}
