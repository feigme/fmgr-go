package enum

import (
	"errors"
	"fmt"
)

type OptionCreateEnum int

const (
	LONG  OptionCreateEnum = 1 // 买
	SHORT OptionCreateEnum = 2 // 卖
)

var optionCreateEnumMap = map[OptionCreateEnum]string{
	SHORT: "卖出",
	LONG:  "买入",
}

func (o OptionCreateEnum) Desc() string {
	str, ok := optionCreateEnumMap[o]
	if ok {
		return str
	}
	return ""
}

// List 列表输出
func OptionCreateEnumList() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range optionCreateEnumMap {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	return km
}

// 是否存在
func GetOptionCreateEnumByKey(key int) (OptionCreateEnum, error) {
	for k := range optionCreateEnumMap {
		if k == OptionCreateEnum(key) {
			return k, nil
		}
	}
	return 0, errors.New("操作定义不存在！")
}
