package models

import "gorm.io/gorm"

type SKU struct {
	gorm.Model
	ProductID       uint             `gorm:"not null" json:"product_id"`
	Price           float64          `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock           int              `gorm:"not null" json:"stock"`
	Code            string           `gorm:"not null;comment:SKU编码"`
	Product         Product          `gorm:"foreignKey:ProductID"`
	AttributeValues []AttributeValue `gorm:"many2many:sku_attribute_values;foreignKey:ID;joinForeignKey:SkuId;otherKey:ID;joinReferences:AttributeValueID" json:"attribute_value,omitempty"`
	Images          []Image          `gorm:"polymorphic:Imageable" json:"images,omitempty"`
}
