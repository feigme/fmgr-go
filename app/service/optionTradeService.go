package service

import (
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/repository"
)

type OptionTradeService struct{}

var OptionTradeSvc = new(OptionTradeService)

func (svc *OptionTradeService) Save(trade *models.OptionTrade) error {
	return repository.OptionTradeRepo.Save(trade)
}

func (svc *OptionTradeService) List(code string) (list []models.OptionTrade) {
	return repository.OptionTradeRepo.List(code)
}

func (svc *OptionTradeService) Get(code string) (*models.OptionTrade, error) {
	return repository.OptionTradeRepo.Get(code)
}

func (svc *OptionTradeService) Close(trade *models.OptionTrade) error {
	return repository.OptionTradeRepo.Update(trade)
}
