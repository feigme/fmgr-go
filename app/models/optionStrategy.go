package models

import (
	"errors"
	"strings"

	"github.com/feigme/fmgr-go/app/enum"
)

type OptionStrategy struct {
	ID
	Timestamps
	Key int
}

func NewOptionStrategy(st enum.OptionStrategyEnum) *OptionStrategy {
	return &OptionStrategy{Key: int(st)}
}

func (st *OptionStrategy) RollPut(option *OptionTrade, closePrice, exerciseDate, sellPrice string) (*OptionTrade, *OptionTrade, error) {
	// 检查是否可进行roll
	if option.Type != "P" && option.Position != "short" {
		return nil, nil, errors.New("sell put才能roll")
	}

	if option.Status != int64(enum.OPTION_STATUS_HAVING) && option.Status != int64(enum.OPTION_STATUS_CLOSE) {
		return nil, nil, errors.New("当前put状态无法roll")
	}

	// 平仓当前put
	option.Close(closePrice)

	// 需要roll的目标put code
	code := strings.ReplaceAll(option.Code, option.ExerciseDate, exerciseDate)

	// 生产option交易对象
	rollOptionTrade, err := NewOptionTrade(code, enum.SHORT, sellPrice)
	if err != nil {
		return nil, nil, err
	}

	return option, rollOptionTrade, nil
}
