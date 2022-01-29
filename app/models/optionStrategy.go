package models

import (
	"github.com/feigme/fmgr-go/app/enum"
)

type OptionStrategy struct {
	ID
	Timestamps
	Code string
}

func NewOptionStrategy(st enum.OptionStrategyEnum) *OptionStrategy {
	return &OptionStrategy{Code: st.Name()}
}
