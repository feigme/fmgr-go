package query

type OpstQuery struct {
	CommonQuery
	Id                int
	Code              string
	StatusList        []int
	StartExerciseDate string
	EndExerciseDate   string
}
