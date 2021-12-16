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
	Position   string `gorm:"type:varchar(32)"` // short or long
	Price      string `gorm:"type:varchar(32)"` // 操作价格
	ClosePrice string `gorm:"type:varchar(32)"` // 平仓价格
	Profit     string `gorm:"type:varchar(32)"` // 收益
	ProfitRate string `gorm:"type:varchar(32)"` // 收益率
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
	trade.Price = fmt.Sprintf("%.2f", pricef)
	trade.Status = int64(enum.OPTION_STATUS_HAVING)
	// 新建option只有买和卖
	if enum.LONG == optCreateEnum {
		trade.Position = "long"
		trade.Count = 1
		// 买，损失权利金，负数
		trade.Premium = fmt.Sprintf("%.2f", -pricef*float64(trade.ContractSize))
	} else if enum.SHORT == optCreateEnum {
		trade.Position = "short"
		trade.Count = -1
		// 卖，获得权利金，正数
		trade.Premium = fmt.Sprintf("%.2f", pricef*float64(trade.ContractSize))
	} else {
		return nil, errors.New("期权操作类型错误！")
	}

	return trade, nil
}

func (trade *OptionTrade) Close(price string) error {

	if int(trade.Status) == int(enum.OPTION_STATUS_CLOSE) || int(trade.Status) == int(enum.OPTION_STATUS_INVALID) {
		return errors.New("该期权已经平仓或者失效！")
	}

	// 平仓
	trade.Status = int64(enum.OPTION_STATUS_CLOSE)

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return errors.New("价格格式错误！")
	}
	trade.ClosePrice = fmt.Sprintf("%.2f", pricef)

	if trade.Position == "short" {
		trade.Profit = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.Price)-pricef)*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.Price)-pricef)/cast.ToFloat64(trade.Price))
	} else if trade.Position == "long" {
		trade.Profit = fmt.Sprintf("%.2f", (-cast.ToFloat64(trade.Price)+pricef)*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (-cast.ToFloat64(trade.Price)+pricef)/cast.ToFloat64(trade.Price))
	}

	return nil

}

func (trade *OptionTrade) Invalid() error {
	if int(trade.Status) == int(enum.OPTION_STATUS_CLOSE) || int(trade.Status) == int(enum.OPTION_STATUS_INVALID) {
		return errors.New("该期权已经平仓或者失效！")
	}

	// 失效
	trade.Status = int64(enum.OPTION_STATUS_INVALID)
	trade.ClosePrice = "0.00"
	trade.Profit = trade.Premium
	if trade.Position == "short" {
		trade.ProfitRate = "1.00"
	} else if trade.Position == "long" {
		trade.ProfitRate = "-1.00"
	}

	return nil
}
