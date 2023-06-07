package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName       string  `gorm:"type:varchar(255)"`
	OriginalPrice     float64 `gorm:"type:decimal(10,2)"`
	CurrentPrice      float64 `gorm:"type:decimal(10,2)"`
	OnSale            bool
	Description       string `gorm:"type:text"`
	VideoURL          string `gorm:"type:varchar(255)"`
	Sales             uint
	ProductAttributes []ProductAttribute
	Categories        []Category `gorm:"many2many:product_categories;"`
	Images            []Image    `gorm:"polymorphic:Imageable;"`
	Reviews           []Review
}

type Attribute struct {
	gorm.Model
	AttributeName     string           `gorm:"type:varchar(255)"`
	AttributeValues   []AttributeValue `gorm:"ForeignKey:AttributeID"`
	ProductAttributes []ProductAttribute
}

type AttributeValue struct {
	gorm.Model
	AttributeID       uint
	Value             string             `gorm:"type:varchar(255)"`
	ProductAttributes []ProductAttribute `gorm:"ForeignKey:ValueID"`
}

type ProductAttribute struct {
	gorm.Model
	ProductID       uint
	AttributeID     uint
	ValueID         uint
	Value           AttributeValue
	PriceAdjustment float64 `gorm:"type:decimal(10,2)"`
	Images          []Image `gorm:"polymorphic:Imageable;"`
}

type Review struct {
	gorm.Model
	ProductID  uint
	UserID     uint
	Rating     int
	ReviewText string  `gorm:"type:text"`
	Images     []Image `gorm:"polymorphic:Imageable;"`
}

type Image struct {
	gorm.Model
	ImageURL      string `gorm:"type:varchar(255)"`
	Link          string
	ImageableID   uint
	ImageableType string `gorm:"polymorphic:Imageable;"`
}

type Category struct {
	gorm.Model
	ParentID      int
	CategoryName  string
	URL           string
	Products      []Product   `gorm:"many2many:product_categories;"`
	Images        []Image     `gorm:"polymorphic:Imageable;"`
	Subcategories []*Category `gorm:"foreignkey:ParentID"`
}

type ProductCategory struct {
	gorm.Model
	ProductID  uint
	CategoryID uint
	ParentID   uint
}

type Currency struct {
	gorm.Model
	Code     string  `gorm:"type:varchar(255)"`
	Exchange float64 `gorm:"type:decimal(10,2)"`
}

type Advertisement struct {
	gorm.Model
	Code   string
	Images []Image `gorm:"polymorphic:Imageable;"`
}

type Collection struct {
	gorm.Model
	Name string          `gorm:"type:varchar(255)"`
	Type string          `gorm:"type:varchar(255)"`
	Code string          `gorm:"type:varchar(255)"`
	Rule json.RawMessage `gorm:"type:json"`
}
