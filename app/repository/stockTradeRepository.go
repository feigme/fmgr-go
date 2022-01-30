package repository

import (
	"context"

	"github.com/feigme/fmgr-go/app/models"
	"gorm.io/gorm"
)

type StockTradeRepository struct {
	db *gorm.DB
}

func NewStockTradeRepo(ctx context.Context) *StockTradeRepository {
	return &StockTradeRepository{
		db: GetDB(ctx),
	}
}

func (repo *StockTradeRepository) Save(trade *models.StockTrade) error {
	return repo.db.Save(trade).Error
}

func (repo *StockTradeRepository) GetOptionCodeByStockCode(code string) string {
	trade := new(models.StockTrade)
	repo.db.Where(" code = ? order by id desc limit 1", code).Find(&trade)
	return trade.OptionCode
}
