package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/spf13/cast"
)

type OptionTrade struct {
	ID
	Timestamps
	Option
	Sid        uint   // optionstrategy关联id
	Position   string `gorm:"type:varchar(32)"` // short or long
	BuyPrice   string `gorm:"type:varchar(32)"` // 买价格
	SellPrice  string `gorm:"type:varchar(32)"` // 卖价格
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

func NewOptionTrade(code string, optCreateEnum enum.OptionCreateEnum, price string) (*OptionTrade, error) {
	trade := new(OptionTrade)
	trade.CreateTime = time.Now()
	trade.UpdateTime = time.Now()

	option, err := NewOption(code)
	if err != nil {
		return nil, err
	}

	trade.Option = *option

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("价格格式错误！")
	}

	trade.Status = int64(enum.OPTION_STATUS_HAVING)
	// 新建option只有买和卖
	if enum.LONG == optCreateEnum {
		trade.Position = "long"
		trade.Count = 1
		// 买，损失权利金，负数
		trade.Premium = fmt.Sprintf("%.2f", -pricef*float64(trade.ContractSize))
		trade.BuyPrice = fmt.Sprintf("%.2f", pricef)
	} else if enum.SHORT == optCreateEnum {
		trade.Position = "short"
		trade.Count = -1
		// 卖，获得权利金，正数
		trade.Premium = fmt.Sprintf("%.2f", pricef*float64(trade.ContractSize))
		trade.SellPrice = fmt.Sprintf("%.2f", pricef)
	} else {
		return nil, errors.New("期权操作类型错误！")
	}

	return trade, nil
}

// 平仓
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

	if trade.Position == "short" {
		trade.BuyPrice = fmt.Sprintf("%.2f", pricef)
		trade.Profit = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.SellPrice)-pricef)*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.SellPrice)-pricef)/cast.ToFloat64(trade.SellPrice))
	} else if trade.Position == "long" {
		trade.SellPrice = fmt.Sprintf("%.2f", pricef)
		trade.Profit = fmt.Sprintf("%.2f", (pricef-cast.ToFloat64(trade.BuyPrice))*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (pricef-cast.ToFloat64(trade.BuyPrice))/cast.ToFloat64(trade.BuyPrice))
	}

	return nil

}

// 失效
func (trade *OptionTrade) Invalid() error {
	if int(trade.Status) == int(enum.OPTION_STATUS_CLOSE) || int(trade.Status) == int(enum.OPTION_STATUS_INVALID) {
		return errors.New("该期权已经平仓或者失效！")
	}

	// 失效
	trade.Status = int64(enum.OPTION_STATUS_INVALID)

	trade.Profit = trade.Premium
	if trade.Position == "short" {
		trade.ProfitRate = "1.00"
		trade.BuyPrice = "0.00"
	} else if trade.Position == "long" {
		trade.ProfitRate = "-1.00"
		trade.SellPrice = "0.00"
	}

	return nil
}

// 行权
func (trade *OptionTrade) Exercise() error {
	if int(trade.Status) == int(enum.OPTION_STATUS_CLOSE) || int(trade.Status) == int(enum.OPTION_STATUS_INVALID) {
		return errors.New("该期权已经平仓或者失效！")
	}

	// 失效
	trade.Status = int64(enum.OPTION_STATUS_EXERCISE)

	trade.Profit = trade.Premium
	if trade.Position == "short" {
		trade.ProfitRate = "1.00"
		trade.BuyPrice = "0.00"
	} else if trade.Position == "long" {
		trade.ProfitRate = "-1.00"
		trade.SellPrice = "0.00"
	}

	return nil
}

// roll put
func (trade *OptionTrade) Roll(closePrice, exerciseDate, sellPrice string) (*OptionTrade, error) {
	// 检查是否可进行roll
	if trade.Type != "P" && trade.Position != "short" {
		return nil, errors.New("sell put才能roll")
	}

	if trade.Status != int64(enum.OPTION_STATUS_HAVING) && trade.Status != int64(enum.OPTION_STATUS_CLOSE) {
		return nil, errors.New("当前put状态无法roll")
	}

	// 平仓当前put
	trade.Close(closePrice)

	// 需要roll的目标put code
	code := strings.ReplaceAll(trade.Code, trade.ExerciseDate, exerciseDate)

	// 生产option交易对象
	rollOptionTrade, err := NewOptionTrade(code, enum.SHORT, sellPrice)
	if err != nil {
		return nil, err
	}

	return rollOptionTrade, nil
}
