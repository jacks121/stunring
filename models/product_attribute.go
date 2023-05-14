package models

import (
	"gorm.io/gorm"
)

type ProductAttribute struct {
	gorm.Model
	ProductID        uint      `gorm:"not null" json:"product_id"`
	AttributeID      uint      `gorm:"not null" json:"attribute_id"`
	AttributeValueID uint      `gorm:"not null" json:"attribute_value_id"`
	Product          Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Attribute        Attribute `gorm:"foreignKey:AttributeID" json:"attribute,omitempty"`
}
