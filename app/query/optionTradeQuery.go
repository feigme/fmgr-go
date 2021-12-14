package query

type OptionTradeQuery struct {
	CommonQuery
	Code              string
	StatusList        []int
	StartExerciseDate string
	EndExerciseDate   string
}
