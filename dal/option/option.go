package option

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"
)

type Option struct {
	Id                   int64     `gorm:"primary_key;AUTO_INCREMENT"`
	CreateTime           time.Time `gorm:"not null"`
	UpdateTime           time.Time `gorm:"not null"`
	Code                 string    `gorm:"type:varchar(64);not null"` // 期权code
	Position             string    `gorm:"type:varchar(64);not null"` // 买方 or 卖方
	StrikePrice          int64     `gorm:"not null"`                  // 行权价
	Premium              int64     `gorm:"not null"`                  // 期权权利金
	Type                 string    `gorm:"type:varchar(64);not null"` // 期权类型：put or call
	UnderlyingSecurities string    `gorm:"type:varchar(64);not null"` // 标的
	ExerciseDate         string    `gorm:"not null"`                  // 行权日
	ContractSize         int64     `gorm:"not null;default:100"`      // 合约数量
	PurchasePrice        string    // 买入价格
	SellPrice            string    // 卖出价格
}

// 指定表名称
func (Option) TableName() string {
	return "option"
}

var (
	reg = regexp.MustCompile(`^([A-Za-z]{1,4})(\d{6})([A-Za-z]{1})(\d{5})$`)
)

func NewOption(code, position, price string) (*Option, error) {
	if code == "" || position == "" || price == "" {
		return nil, errors.New("新建期权：参数不全!")
	}

	// code转成大写
	code = strings.ToUpper(code)
	// position转小写
	position = strings.ToLower(position)

	if position != "short" && position != "long" {
		return nil, errors.New(fmt.Sprintf("新建期权：参数错误，position值必须为short or long，当前为: %s", position))
	}

	regResult := reg.FindStringSubmatch(code)
	option := &Option{
		Code:                 code,
		Position:             position,
		StrikePrice:          cast.ToInt64(regResult[4]),
		Type:                 regResult[3],
		UnderlyingSecurities: regResult[1],
		ExerciseDate:         regResult[2],
		ContractSize:         100,
	}

	if position == "long" {
		option.PurchasePrice = price
	}
	if position == "short" {
		option.SellPrice = price
	}

	return option, nil
}
