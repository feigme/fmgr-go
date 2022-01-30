package query

type StockQuery struct {
	CommonQuery
	Id         int
	Code       string
	OptionCode string
	StatusList []int
	CostPrice  string
}
