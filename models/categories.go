package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ParentID *uint      `gorm:"default:null" json:"parent_id,omitempty"`
	Name     string     `gorm:"not null" json:"name"`
	URL      string     `gorm:"not null" json:"url"`
	Children []Category `gorm:"foreignkey:ParentID" json:"children,omitempty"`
	Products []Product  `gorm:"many2many:product_category;" json:"products,omitempty"`
}
