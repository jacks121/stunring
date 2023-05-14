package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string          `gorm:"not null" json:"name"`
	Description  string          `gorm:"not null" json:"description"`
	Price        float64         `gorm:"not null" json:"price"`
	RegularPrice float64         `gorm:"not null;" json:"regular_price"`
	Status       int             `gorm:"not null;default:0;" json:"-"`
	StatusText   string          `gorm:"-" json:"status_text"`
	Categories   []Category      `gorm:"many2many:product_category;" json:"categories,omitempty"`
	Images       []Image         `gorm:"polymorphic:Imageable;"`
	Attributes   []Attribute     `gorm:"many2many:product_attributes" json:"attribute,omitempty"`
	SKUs         []SKU           `gorm:"foreignKey:ProductID" json:"sku,omitempty"`
	Reviews      []ProductReview `gorm:"foreignKey:ProductID" json:"reviews,omitempty"`
}

func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	if p.Status == 0 {
		p.StatusText = "下架"
	} else if p.Status == 1 {
		p.StatusText = "上架"
	}
	return
}
