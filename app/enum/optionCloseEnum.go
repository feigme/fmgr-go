package enum

import (
	"errors"
	"fmt"
)

type OptionCloseEnum string

const (
	OPTION_OPS_CLOSE    OptionCloseEnum = "close"    // 平仓
	OPTION_OPS_INVALID  OptionCloseEnum = "invalid"  // 失效
	OPTION_OPS_EXERCISE OptionCloseEnum = "exercise" // 行权
)

var optionCloseEnumMap = map[OptionCloseEnum]string{
	OPTION_OPS_CLOSE:    "平仓",
	OPTION_OPS_INVALID:  "失效",
	OPTION_OPS_EXERCISE: "行权",
}

func (o OptionCloseEnum) Desc() string {
	p, ok := optionCloseEnumMap[o]
	if ok {
		return p
	}
	return ""
}

func (o OptionCloseEnum) Name() string {
	return string(o)
}

// List 列表输出
func OptionCloseEnumList() []EnumItem {
	km := make([]EnumItem, 0)
	for k, v := range optionCloseEnumMap {
		km = append(km, EnumItem{Name: k.Name(), Desc: v})
	}
	return km
}

// 是否存在
func GetOptionCloseEnumByName(name string) (OptionCloseEnum, error) {
	for k := range optionCloseEnumMap {
		if name == k.Name() {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprintf("枚举不存在! name: %s", name))
}
