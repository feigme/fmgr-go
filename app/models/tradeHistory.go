package models

import "time"

type TradeHistory struct {
	ID
	Timestamps
	Code      string    `gorm:"type:varchar(16);not null"` // 代码
	Name      string    `gorm:"type:varchar(64);not null"` // 名称
	Direction string    `gorm:"type:varchar(8);not null"`  // 方向
	Count     string    `gorm:"not null"`                  // 成交数
	Price     string    `gorm:"not null"`                  // 成交价格
	Amount    string    `gorm:"not null"`                  // 金额
	TradeTime time.Time `gorm:"not null"`                  // 成交时间
}

func NewTradeHistory() *TradeHistory {
	obj := new(TradeHistory)
	obj.CreateTime = time.Now()
	obj.UpdateTime = time.Now()
	return obj
}
