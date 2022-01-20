package enum

type OptionStatusEnum int

const (
	OPTION_STATUS_HAVING   OptionStatusEnum = 1  // 持仓
	OPTION_STATUS_INVALID  OptionStatusEnum = -1 // 失效
	OPTION_STATUS_CLOSE    OptionStatusEnum = 2  // 平仓
	OPTION_STATUS_EXERCISE OptionStatusEnum = 3  // 行权
	OPTION_STATUS_ROLL     OptionStatusEnum = 4  // roll
)

var optionStatusEnumMap = map[OptionStatusEnum]string{
	OPTION_STATUS_HAVING:   "持仓",
	OPTION_STATUS_INVALID:  "失效",
	OPTION_STATUS_CLOSE:    "平仓",
	OPTION_STATUS_EXERCISE: "行权",
	OPTION_STATUS_ROLL:     "roll",
}

func (o OptionStatusEnum) Desc() string {
	str, ok := optionStatusEnumMap[o]
	if ok {
		return str
	}
	return ""
}
