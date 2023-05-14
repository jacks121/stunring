package models

type Image struct {
	ID            uint `gorm:"primary_key"`
	URL           string
	ImageableID   uint
	ImageableType string `gorm:"polymorphic:Imageable"`
}
