package models

import "time"

type Currency struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"type:char(3);uniqueIndex;not null"`
	Name      string `gorm:"type:varchar(50);not null"`
	Symbol    string `gorm:"type:varchar(5);not null"`
	CreatedAt time.Time
}

var DefaultCurrencies = []Currency{
	{Code: "USD", Name: "US Dollar", Symbol: "$"},
	{Code: "EUR", Name: "Euro", Symbol: "€"},
	{Code: "GBP", Name: "British Pound", Symbol: "£"},
	{Code: "JPY", Name: "Japanese Yen", Symbol: "¥"},
	{Code: "SGD", Name: "Singapore Dollar", Symbol: "S$"},
	{Code: "AUD", Name: "Australian Dollar", Symbol: "A$"},
	{Code: "KRW", Name: "South Korean Won", Symbol: "₩"},
	{Code: "IDR", Name: "Indonesian Rupiah", Symbol: "Rp"},
	{Code: "CNY", Name: "Chinese Yuan", Symbol: "¥"},
}
