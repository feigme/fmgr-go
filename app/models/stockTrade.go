package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/feigme/fmgr-go/app/enum"
)

type StockTrade struct {
	ID
	Timestamps
	Stock
	OptionCode string `gorm:"type:varchar(8)"`           // 期权code
	CostPrice  string `gorm:"type:varchar(32);not null"` // 成本价
	SellPrice  string `gorm:"type:varchar(32);not null"` // 卖出价
	Count      int64  `gorm:"not null"`                  // 持仓数量
	Profit     string `gorm:"type:varchar(64)"`          // 收益
	Status     string `gorm:"type:varchar(16)"`          // 状态
	Market     string `gorm:"type:varchar(8)"`           // 股票市场 HK/US
}

// 自定义表名
func (StockTrade) TableName() string {
	return "stock_trade"
}

func NewStockTrade(market, code, optionCode, price string, count int) (*StockTrade, error) {
	stock := new(StockTrade)
	stock.CreateTime = time.Now()
	stock.UpdateTime = time.Now()
	stock.Status = enum.Stock_STATUS_HAVING.Name()

	market = strings.ToUpper(market)
	if market != "HK" && market != "US" {
		return nil, errors.New("只不支持US、HK! ")
	}
	stock.Market = market
	stock.Code = code
	stock.OptionCode = optionCode

	_, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("价格格式错误! ")
	}
	stock.CostPrice = price

	stock.Count = int64(count)

	return stock, nil
}
