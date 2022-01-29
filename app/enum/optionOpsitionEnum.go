package enum

import (
	"errors"
	"fmt"
)

type OptionPositionEnum string // 期权的买方、卖方

const (
	Option_Position_Seller OptionPositionEnum = "seller" // 卖方
	Option_Position_Buyer  OptionPositionEnum = "buyer"  // 买方
)

var optionPositionEnumMap = map[OptionPositionEnum]string{
	Option_Position_Seller: "卖方",
	Option_Position_Buyer:  "买方",
}

func (o OptionPositionEnum) Name() string {
	return string(o)
}

func (o OptionPositionEnum) Desc() string {
	p, ok := optionPositionEnumMap[o]
	if ok {
		return p
	}
	return ""
}

// List 列表输出
func OptionPositionEnumList() []EnumItem {
	km := make([]EnumItem, 0)
	for k, v := range optionPositionEnumMap {
		km = append(km, EnumItem{Name: k.Name(), Desc: v})
	}
	return km
}

// 是否存在
func GetOptionPositionEnumByName(name string) (OptionPositionEnum, error) {
	for k := range optionPositionEnumMap {
		if name == k.Name() {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprintf("枚举不存在! name: %s", name))
}
