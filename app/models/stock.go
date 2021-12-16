package models

type Stock struct {
	Code string `gorm:"type:varchar(64);not null"` // 股票code
}
