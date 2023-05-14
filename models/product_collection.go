package models

import "gorm.io/gorm"

type ProductCollection struct {
	gorm.Model
	CollectionName string `gorm:"type:varchar(50);not null" json:"collection_name"`
	Description    string `gorm:"type:varchar(255);not null" json:"description"`
	Conditions     string `gorm:"type:json;not null" json:"conditions"`
}
