package service

import (
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/repository"
)

type OptionTradeService struct{}

var OptionTradeSvc = new(OptionTradeService)

func (o *OptionTradeService) Save(trade *models.OptionTrade) error {
	err := repository.OptionTradeRepo.Save(trade)

	return err
}
