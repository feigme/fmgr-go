package service

import (
	"context"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/repository"
)

type StockTradeService struct {
}

var StockTradeSvc = new(StockTradeService)

func (svc *StockTradeService) GetOptionCode(code string) string {
	return repository.NewStockTradeRepo(context.Background()).GetOptionCodeByStockCode(code)
}

func (svc *StockTradeService) Save(trade *models.StockTrade) error {
	return repository.NewStockTradeRepo(context.Background()).Save(trade)
}
