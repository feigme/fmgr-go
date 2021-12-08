package enum

import (
	"errors"
	"fmt"
)

type OptionCloseEnum int

const (
	CLOSE   OptionCloseEnum = 3  // 平仓
	INVALID OptionCloseEnum = -1 // 失效
)

var optionCloseEnumMap = map[OptionCloseEnum]string{
	CLOSE:   "平仓",
	INVALID: "失效",
}

func (o OptionCloseEnum) Desc() string {
	str, ok := optionCloseEnumMap[o]
	if ok {
		return str
	}
	return ""
}

// List 列表输出
func OptionCloseEnumList() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range optionCloseEnumMap {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	return km
}

// 是否存在
func GetOptionCloseEnumByKey(key int) (OptionCloseEnum, error) {
	for k := range optionCloseEnumMap {
		if k == OptionCloseEnum(key) {
			return k, nil
		}
	}
	return 0, errors.New("操作定义不存在！")
}
