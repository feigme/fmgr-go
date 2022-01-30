package query

type OptionQuery struct {
	CommonQuery
	Id                int
	Code              string
	StatusList        []int
	StartExerciseDate string
	EndExerciseDate   string
	Position          string
}
