package models

type Currency struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	CurrencyCode string  `gorm:"type:varchar(10);not null" json:"currency_code"`
	CurrencyName string  `gorm:"type:varchar(50);not null" json:"currency_name"`
	ExchangeRate float64 `gorm:"type:decimal(10,2);default:0.00" json:"exchange_rate"`
}
