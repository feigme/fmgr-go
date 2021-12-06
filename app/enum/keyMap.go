package enum

// KeyMap 输出模型
type KeyMap struct {
	Key string `json:"k"`
	Val int    `json:"v"`
}

// KeyMapSlice
type KeyMapSlice []KeyMap

func (s KeyMapSlice) Len() int {
	return len(s)
}
