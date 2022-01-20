package enum

import (
	"errors"
	"sort"
)

//期权策略
type OptionStrategyEnum int

type OptionStrategyInfo struct {
	Key  int
	Code string
	Desc string
}

const (
	OST_COVERED_CALL OptionStrategyEnum = 1
	OST_ROLLING_PUT  OptionStrategyEnum = 2
)

var optionStrategyMap = map[OptionStrategyEnum]OptionStrategyInfo{
	OST_COVERED_CALL: {
		Key:  1,
		Code: "covered call",
		Desc: "持有一手正股，同时卖出1张call",
	},
	OST_ROLLING_PUT: {
		Key:  2,
		Code: "rolling put",
		Desc: "平仓当前put，并卖出1张行权价相同、行权日比较靠后的put",
	},
}

func (o OptionStrategyEnum) getCode() string {
	info, ok := optionStrategyMap[o]
	if ok {
		return info.Code
	}
	return ""
}

func (o OptionStrategyEnum) getDesc() string {
	info, ok := optionStrategyMap[o]
	if ok {
		return info.Desc
	}
	return ""
}

// List 列表输出
func OptionStEnumList() []OptionStrategyInfo {
	// 所有key
	var keys []int
	for k := range optionStrategyMap {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	km := make([]OptionStrategyInfo, 0)
	for _, k := range keys {
		km = append(km, optionStrategyMap[OptionStrategyEnum(k)])
	}
	return km
}

// 是否存在
func GetOptionStEnumByKey(key int) (OptionStrategyEnum, error) {
	for k := range optionStrategyMap {
		if k == OptionStrategyEnum(key) {
			return k, nil
		}
	}
	return 0, errors.New("操作定义不存在！")
}
