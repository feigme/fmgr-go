package models

import (
	"errors"
	"regexp"
	"strings"
)

type Option struct {
	Code         string `gorm:"type:varchar(64);not null"` // 期权code
	Stock        string `gorm:"type:varchar(64);not null"` // 标的
	ExerciseDate string `gorm:"not null"`                  // 行权日
	Type         string `gorm:"type:varchar(64);not null"` // 期权类型：put or call
	StrikePrice  string `gorm:"not null"`                  // 行权价
	ContractSize int64  `gorm:"not null"`                  // 合约数量
}

var (
	reg = regexp.MustCompile(`^([A-Za-z]{1,4})(\d{6})([pPcC]{1})(\d{5})$`)
)

func NewOption(code string) (*Option, error) {
	if code == "" {
		return nil, errors.New("期权code不能为空！")
	}

	// code转成大写
	code = strings.ToUpper(code)
	err := reg.MatchString(code)
	if !err {
		return nil, errors.New("期权code格式错误！")
	}

	regResult := reg.FindStringSubmatch(code)
	option := &Option{
		Code:         code,
		Stock:        regResult[1],
		ExerciseDate: regResult[2],
		Type:         regResult[3],
		StrikePrice:  regResult[4],
		ContractSize: 100,
	}

	return option, nil
}
