package enum

import (
	"errors"
	"fmt"
)

type StockStatusEnum string

const (
	Stock_STATUS_HAVING StockStatusEnum = "having" // 持仓
	Stock_STATUS_LOCK   StockStatusEnum = "lock"   // 锁定
	Stock_STATUS_CLOSE  StockStatusEnum = "close"  // 平仓
)

var stockStatusEnumMap = map[StockStatusEnum]string{
	Stock_STATUS_HAVING: "持仓",
	Stock_STATUS_LOCK:   "锁定",
	Stock_STATUS_CLOSE:  "平仓",
}

func (o StockStatusEnum) Desc() string {
	p, ok := stockStatusEnumMap[o]
	if ok {
		return p
	}
	return ""
}

func (o StockStatusEnum) Name() string {
	return string(o)
}

// List 列表输出
func StockStatusEnumList() []EnumItem {
	km := make([]EnumItem, 0)
	for k, v := range stockStatusEnumMap {
		km = append(km, EnumItem{Name: k.Name(), Desc: v})
	}
	return km
}

// 是否存在
func GetStockStatusEnumByName(name string) (StockStatusEnum, error) {
	for k := range stockStatusEnumMap {
		if k.Name() == name {
			return k, nil
		}
	}
	return "", errors.New(fmt.Sprintf("枚举不存在! name: %s", name))
}
