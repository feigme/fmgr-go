package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/spf13/cast"
)

type OptionTrade struct {
	ID
	Timestamps
	Option
	Position   string `gorm:"type:varchar(64)"` // short or long
	Price      string `gorm:"type:varchar(64)"` // 操作价格
	ClosePrice string `gorm:"type:varchar(64)"` // 平仓价格
	Profit     string `gorm:"type:varchar(64)"` // 收益
	Count      int64  `gorm:"not null"`         // 数量
	Premium    string `gorm:"not null"`         // 期权权利金
	Status     int64  `gorm:"not null"`         // 状态
}

// 自定义表名
func (OptionTrade) TableName() string {
	return "option_trade"
}

func NewOptionTrade(option *Option, optCreateEnum enum.OptionCreateEnum, price string) (*OptionTrade, error) {
	trade := new(OptionTrade)
	trade.CreateTime = time.Now()
	trade.UpdateTime = time.Now()

	if option == nil {
		return nil, errors.New("缺少期权信息！")
	}
	trade.Option = *option

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("价格格式错误！")
	}
	trade.Price = price

	// 新建option只有买和卖
	if enum.LONG == optCreateEnum {
		trade.Position = "long"
		trade.Count = 1
	} else if enum.SHORT == optCreateEnum {
		trade.Position = "short"
		trade.Count = -1
	} else {
		return nil, errors.New("期权操作类型错误！")
	}
	trade.Status = int64(optCreateEnum)
	trade.Premium = fmt.Sprintf("%.0f", pricef*float64(trade.ContractSize)*float64(trade.Count))

	return trade, nil
}

func (trade *OptionTrade) Close(optCloseEnum enum.OptionCloseEnum, price string) error {

	if optCloseEnum != enum.CLOSE {
		return errors.New("操作类型错误！")
	}
	// 平仓
	trade.Status = int64(enum.CLOSE)

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return errors.New("价格格式错误！")
	}
	trade.ClosePrice = price

	cnt := -float64(trade.Count)
	trade.Profit = fmt.Sprintf("%.0f", cast.ToFloat64(trade.Premium)+pricef*cnt)

	return nil

}

func (trade *OptionTrade) Invalid(optCloseEnum enum.OptionCloseEnum) error {
	if optCloseEnum != enum.INVALID {
		return errors.New("操作类型错误！")
	}

	// 失效
	trade.Status = int64(enum.INVALID)
	trade.ClosePrice = "0"
	trade.Profit = trade.Premium
	return nil
}
