package models

import (
	"time"

	"github.com/feigme/fmgr-go/app/enum"
)

type OptionStrategy struct {
	ID
	Timestamps
	Name      string `gorm:"type:varchar(128)"`
	Namespace string `gorm:"type:varchar(16)"`
	Labels    string `gorm:"type:varchar(512)"`
	Code      string `gorm:"type:varchar(64)"`
}

func NewOptionStrategy(st enum.OptionStrategyEnum) *OptionStrategy {
	strategy := new(OptionStrategy)
	strategy.Code = st.Name()
	strategy.CreateTime = time.Now()
	strategy.UpdateTime = time.Now()

	return strategy
}
