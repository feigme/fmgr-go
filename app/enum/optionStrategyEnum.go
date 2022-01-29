package enum

import (
	"errors"
	"fmt"
)

//期权策略
type OptionStrategyEnum string

const (
	OST_Covered_Call     OptionStrategyEnum = "covered call"
	OST_Collar           OptionStrategyEnum = "collar"
	OST_Naked_Short_Put  OptionStrategyEnum = "naked short put"
	OST_Naked_Short_Call OptionStrategyEnum = "naked short call"
)

var optionStrategyMap = map[OptionStrategyEnum]string{
	OST_Covered_Call:     "持有一手正股，同时卖出1张价外call",
	OST_Collar:           "持有一手正股，同时卖出1张价外call，买入1张价外put",
	OST_Naked_Short_Put:  "裸卖1张价外put",
	OST_Naked_Short_Call: "裸卖1张价外call",
}

func (o OptionStrategyEnum) Desc() string {
	p, ok := optionStrategyMap[o]
	if ok {
		return p
	}
	return ""
}

func (o OptionStrategyEnum) Name() string {
	return string(o)
}

// List 列表输出
func OptionStrategyEnumList() []EnumItem {
	km := make([]EnumItem, 0)
	for k, v := range optionStrategyMap {
		km = append(km, EnumItem{Name: k.Name(), Desc: v})
	}
	return km
}

// 是否存在
func GetOptionStEnumByName(name string) (OptionStrategyEnum, error) {
	for k := range optionStrategyMap {
		if k.Name() == name {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprintf("枚举不存在! name: %s", name))
}
