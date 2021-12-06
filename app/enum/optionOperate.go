package enum

import (
	"errors"
	"fmt"
)

type OptionOperate int

const (
	SHORT   OptionOperate = 1  // 卖
	LONG    OptionOperate = 2  // 买
	CLOSE   OptionOperate = 3  // 平仓
	INVALID OptionOperate = -1 // 失效
)

var optionOperateMap = map[OptionOperate]string{
	SHORT:   "卖出",
	LONG:    "买入",
	CLOSE:   "平仓",
	INVALID: "失效",
}

func (o OptionOperate) Desc() string {
	str, ok := optionOperateMap[o]
	if ok {
		return str
	}
	return ""
}

// List 列表输出
func OptionOperateList() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range optionOperateMap {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	return km
}

// 是否存在
func GetOptionOperateByKey(key int) (OptionOperate, error) {
	for k := range optionOperateMap {
		if k == OptionOperate(key) {
			return k, nil
		}
	}
	return 0, errors.New("操作定义不存在！")
}
