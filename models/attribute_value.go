package models

import (
	"gorm.io/gorm"
)

type AttributeValue struct {
	gorm.Model
	AttributeID uint      `gorm:"not null" json:"attribute_id"`
	Value       string    `gorm:"not null" json:"value"`
	Attribute   Attribute `gorm:"foreignKey:AttributeID" json:"attribute,omitempty"`
	SKUs        []SKU     `gorm:"many2many:sku_attribute_values;foreignKey:ID;joinForeignKey:AttributeValueID;otherKey:ID;joinReferences:SkuId" json:"sku,omitempty"`
}
