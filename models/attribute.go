package models

import (
	"gorm.io/gorm"
)

type Attribute struct {
	gorm.Model
	Name            string           `gorm:"not null" json:"name"`
	AttributeValues []AttributeValue `gorm:"foreignKey:AttributeID" json:"attribute_value,omitempty"`
	Products        []Product        `gorm:"many2many:product_attributes" json:"products,omitempty"`
}
