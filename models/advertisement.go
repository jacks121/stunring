package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model
	Position string `gorm:"type:varchar(50);not null" json:"position"`
	ImageURL string `gorm:"type:varchar(255);not null" json:"image_url"`
	ClickURL string `gorm:"type:varchar(255);not null" json:"click_url"`
}
