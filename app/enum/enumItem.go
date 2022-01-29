package enum

// enum 输出模型
type EnumItem struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type EnumCellSlice []EnumItem

func (s EnumCellSlice) Len() int {
	return len(s)
}
