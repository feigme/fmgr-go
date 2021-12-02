package repository

import (
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/global"
)

type OptionRepo struct{}

func (o *OptionRepo) Save(code, position, price string) (*models.Option, error) {
	option, err := models.NewOption(code, position, price)
	if err != nil {
		return nil, err
	}
	return option, global.App.DB.Save(option).Error
}

var optionRepo = new(OptionRepo)
