package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/spf13/cast"
)

type OptionTrade struct {
	ID
	Timestamps
	Option
	Position  string `gorm:"type:varchar(64)"` // short or long
	BuyPrice  string `gorm:"type:varchar(64)"` // 买入价格
	SellPrice string `gorm:"type:varchar(64)"` // 卖出价格
	Count     int64  `gorm:"not null"`         // 数量
	Premium   string `gorm:"not null"`         // 期权权利金
	Status    int64  `gorm:"not null"`         // 状态
}

// 自定义表名
func (OptionTrade) TableName() string {
	return "option_trade"
}

func NewOptionTrade(option *Option, operate enum.OptionOperate, price string) (*OptionTrade, error) {
	trade := &OptionTrade{
		Status: int64(operate),
		Count:  100,
	}
	trade.CreateTime = time.Now()
	trade.UpdateTime = time.Now()

	trade.Option = *option

	amount, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("价格格式不对！")
	}

	// 新建option只有买和卖
	if enum.LONG == operate {
		trade.BuyPrice = price
		trade.Position = "long"
		trade.Premium = cast.ToString(amount * float64(trade.Count))
	} else if enum.SHORT == operate {
		trade.SellPrice = price
		trade.Position = "short"
		trade.Premium = cast.ToString(amount * float64(trade.Count))
	} else {
		return nil, errors.New("期权操作类型错误！")
	}

	return trade, nil
}
