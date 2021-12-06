package enum

//期权策略
type OptionStrategy int

const (
	COVERED_CALL OptionStrategy = 1
)

var optionStrategyMap = map[OptionStrategy]string{
	COVERED_CALL: "covered call",
}

func (o OptionStrategy) getDesc() string {
	str, ok := optionStrategyMap[o]
	if ok {
		return str
	}
	return ""
}
