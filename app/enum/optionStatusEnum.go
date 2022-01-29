package enum

import (
	"errors"
	"fmt"
)

type OptionStatusEnum string

const (
	OPTION_STATUS_HAVING   OptionStatusEnum = "having"   // 持仓
	OPTION_STATUS_INVALID  OptionStatusEnum = "invalid"  // 失效
	OPTION_STATUS_CLOSE    OptionStatusEnum = "close"    // 平仓
	OPTION_STATUS_EXERCISE OptionStatusEnum = "exercise" // 行权
)

var optionStatusEnumMap = map[OptionStatusEnum]string{
	OPTION_STATUS_HAVING:   "持仓",
	OPTION_STATUS_INVALID:  "失效",
	OPTION_STATUS_CLOSE:    "平仓",
	OPTION_STATUS_EXERCISE: "行权",
}

func (o OptionStatusEnum) Desc() string {
	p, ok := optionStatusEnumMap[o]
	if ok {
		return p
	}
	return ""
}

func (o OptionStatusEnum) Name() string {
	return string(o)
}

// List 列表输出
func OptionStatusEnumList() []EnumItem {
	km := make([]EnumItem, 0)
	for k, v := range optionStatusEnumMap {
		km = append(km, EnumItem{Name: k.Name(), Desc: v})
	}
	return km
}

// 是否存在
func GetOptionStatusEnumByName(name string) (OptionStatusEnum, error) {
	for k := range optionStatusEnumMap {
		if k.Name() == name {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprintf("枚举不存在! name: %s", name))
}
