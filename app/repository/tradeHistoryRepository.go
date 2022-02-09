package repository

import (
	"context"
	"errors"

	"github.com/feigme/fmgr-go/app/models"
	"gorm.io/gorm"
)

type TradeHistoryRepository struct {
	db *gorm.DB
}

func NewTradeHistoryRepo(ctx context.Context) *TradeHistoryRepository {
	return &TradeHistoryRepository{
		db: GetDB(ctx),
	}
}

func (repo *TradeHistoryRepository) Save(trade *models.TradeHistory) error {
	return repo.db.Save(trade).Error
}

func (repo *TradeHistoryRepository) IsExisted(trade *models.TradeHistory) bool {
	err := repo.db.Where(" code = ? and count = ? and price = ? and trade_time = ?", trade.Code, trade.Count, trade.Price, trade.TradeTime).First(&trade).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
