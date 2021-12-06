package repository

import (
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/global"
)

type OptionTradeRepository struct{}

var OptionTradeRepo = new(OptionTradeRepository)

func (o *OptionTradeRepository) Save(trade *models.OptionTrade) error {
	return global.App.DB.Save(trade).Error
}
