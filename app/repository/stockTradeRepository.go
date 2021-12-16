package repository

import (
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/global"
)

type StockTradeRepository struct{}

func (repo *StockTradeRepository) Save(trade *models.StockTrade) error {
	return global.App.DB.Save(trade).Error
}
