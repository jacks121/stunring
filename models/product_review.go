package models

import "gorm.io/gorm"

type ProductReview struct {
	gorm.Model
	ProductID uint
	UserID    uint
	Rating    uint
	Nickname  string
	Summary   string
	Review    string
	Images    []*Image `gorm:"polymorphic:Imageable;polymorphicValue:product_review"`
}
