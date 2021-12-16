package models

type StockTrade struct {
	ID
	Timestamps
	Stock
	CostPrice string `gorm:"type:varchar(32);not null"` // 成本价
	SellPrice string `gorm:"type:varchar(32);not null"` // 卖出价
	Count     int64  `gorm:"not null"`                  // 持仓数量
	Profit    string `gorm:"type:varchar(64)"`          // 收益
}
