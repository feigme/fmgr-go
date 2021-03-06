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
	Sid        uint      // optionstrategy关联id
	Position   string    `gorm:"type:varchar(32)"` // seller or buyer
	BuyPrice   string    `gorm:"type:varchar(32)"` // 买价格
	SellPrice  string    `gorm:"type:varchar(32)"` // 卖价格
	Profit     string    `gorm:"type:varchar(32)"` // 收益
	ProfitRate string    `gorm:"type:varchar(32)"` // 收益率
	Count      int64     `gorm:"not null"`         // 数量
	Premium    string    `gorm:"not null"`         // 期权权利金
	Status     string    `gorm:"not null"`         // 状态
	Market     string    `gorm:"type:varchar(8)"`  // 股票市场 HK/US
	OpenTime   time.Time // 开仓时间
	CloseTime  time.Time // 平仓时间
}

// 自定义表名
func (OptionTrade) TableName() string {
	return "option_trade"
}

func NewOptionTrade(market, code string, optPositionEnum enum.OptionPositionEnum, price string) (*OptionTrade, error) {
	trade := new(OptionTrade)
	trade.CreateTime = time.Now()
	trade.UpdateTime = time.Now()

	option, err := NewOption(code)
	if err != nil {
		return nil, err
	}
	trade.Option = *option

	market = strings.ToUpper(market)
	if market != "HK" && market != "US" {
		return nil, errors.New("只不支持US、HK! ")
	}
	trade.Market = market
	if trade.Market == "US" {
		trade.ContractSize = 100 // 美股固定是100
	}

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, errors.New("价格格式错误! ")
	}

	trade.Status = enum.OPTION_STATUS_HAVING.Name()
	trade.Position = optPositionEnum.Name()
	// 新建option只有买和卖
	if enum.Option_Position_Buyer == optPositionEnum {
		trade.Count = 1
		// 买，损失权利金，负数
		trade.Premium = fmt.Sprintf("%.2f", -pricef*float64(trade.ContractSize))
		trade.BuyPrice = fmt.Sprintf("%.2f", pricef)
	} else if enum.Option_Position_Seller == optPositionEnum {
		trade.Count = -1
		// 卖，获得权利金，正数
		trade.Premium = fmt.Sprintf("%.2f", pricef*float64(trade.ContractSize))
		trade.SellPrice = fmt.Sprintf("%.2f", pricef)
	} else {
		return nil, errors.New("期权操作类型错误! ")
	}

	return trade, nil
}

// 平仓
func (trade *OptionTrade) Close(price string) error {

	if trade.Status == enum.OPTION_STATUS_CLOSE.Name() || trade.Status == enum.OPTION_STATUS_INVALID.Name() {
		return errors.New("该期权已经平仓或者失效! ")
	}

	// 平仓
	trade.Status = enum.OPTION_STATUS_CLOSE.Name()

	pricef, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return errors.New("价格格式错误! ")
	}

	if trade.Position == enum.Option_Position_Seller.Name() {
		trade.BuyPrice = fmt.Sprintf("%.2f", pricef)
		trade.Profit = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.SellPrice)-pricef)*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (cast.ToFloat64(trade.SellPrice)-pricef)/cast.ToFloat64(trade.SellPrice))
	} else if trade.Position == enum.Option_Position_Buyer.Name() {
		trade.SellPrice = fmt.Sprintf("%.2f", pricef)
		trade.Profit = fmt.Sprintf("%.2f", (pricef-cast.ToFloat64(trade.BuyPrice))*float64(trade.ContractSize))
		trade.ProfitRate = fmt.Sprintf("%.2f", (pricef-cast.ToFloat64(trade.BuyPrice))/cast.ToFloat64(trade.BuyPrice))
	}

	if trade.CloseTime.IsZero() {
		trade.CloseTime = time.Now()
	}

	return nil

}

// 失效
func (trade *OptionTrade) Invalid() error {
	if trade.Status == enum.OPTION_STATUS_CLOSE.Name() || trade.Status == enum.OPTION_STATUS_INVALID.Name() {
		return errors.New("该期权已经平仓或者失效! ")
	}

	// 失效
	trade.Status = enum.OPTION_STATUS_INVALID.Name()

	trade.Profit = trade.Premium
	if trade.Position == enum.Option_Position_Seller.Name() {
		trade.ProfitRate = "1.00"
		trade.BuyPrice = "0.00"
	} else if trade.Position == enum.Option_Position_Buyer.Name() {
		trade.ProfitRate = "-1.00"
		trade.SellPrice = "0.00"
	}

	if trade.CloseTime.IsZero() {
		trade.CloseTime = time.Now()
	}
	return nil
}

// 行权
func (trade *OptionTrade) Exercise() error {
	if trade.Status == enum.OPTION_STATUS_CLOSE.Name() || trade.Status == enum.OPTION_STATUS_INVALID.Name() {
		return errors.New("该期权已经平仓或者失效! ")
	}

	// 失效
	trade.Status = enum.OPTION_STATUS_EXERCISE.Name()

	trade.Profit = trade.Premium
	if trade.Position == enum.Option_Position_Seller.Name() {
		trade.ProfitRate = "1.00"
		trade.BuyPrice = "0.00"
	} else if trade.Position == enum.Option_Position_Buyer.Name() {
		trade.ProfitRate = "-1.00"
		trade.SellPrice = "0.00"
	}

	if trade.CloseTime.IsZero() {
		trade.CloseTime = time.Now()
	}

	return nil
}

// roll put
func (trade *OptionTrade) Roll(closePrice, exerciseDate, sellPrice string) (*OptionTrade, error) {
	// 检查是否可进行roll
	if trade.Type != "P" && trade.Position != "short" {
		return nil, errors.New("sell put才能roll")
	}

	if trade.Status != enum.OPTION_STATUS_HAVING.Name() && trade.Status != enum.OPTION_STATUS_CLOSE.Name() {
		return nil, errors.New("当前put状态无法roll")
	}

	// 平仓当前put
	trade.Close(closePrice)

	// 需要roll的目标put code
	code := strings.ReplaceAll(trade.Code, trade.ExerciseDate, exerciseDate)

	// 生产option交易对象
	rollOptionTrade, err := NewOptionTrade(trade.Market, code, enum.Option_Position_Seller, sellPrice)
	if err != nil {
		return nil, err
	}

	return rollOptionTrade, nil
}
