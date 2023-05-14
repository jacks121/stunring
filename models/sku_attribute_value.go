package models

import "gorm.io/gorm"

type SKUAttributeValue struct {
	gorm.Model
	SKUID            uint           `gorm:"column:sku_id;not null" json:"sku_id"`
	AttributeValueID uint           `gorm:"not null" json:"attribute_value_id"`
	SKU              SKU            `gorm:"foreignKey:SkuId" json:"sku,omitempty"`
	AttributeValue   AttributeValue `gorm:"foreignKey:AttributeValueID" json:"attribute_value,omitempty"`
}
