package service

import (
	"errors"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/repository"
)

type OptionTradeService struct{}

var OptionTradeSvc = new(OptionTradeService)

func (svc *OptionTradeService) Save(trade *models.OptionTrade) error {
	return repository.OptionTradeRepo.Save(trade)
}

func (svc *OptionTradeService) List(query *query.OptionTradeQuery) (list []models.OptionTrade) {
	return repository.OptionTradeRepo.List(query)
}

func (svc *OptionTradeService) Get(code string) (*models.OptionTrade, error) {
	return repository.OptionTradeRepo.Get(code)
}

func (svc *OptionTradeService) Close(trade *models.OptionTrade) error {
	return repository.OptionTradeRepo.Update(trade)
}

func (svc *OptionTradeService) Delete(query *query.OptionTradeQuery) error {
	list := repository.OptionTradeRepo.List(query)
	if len(list) == 0 {
		return errors.New("记录为空，请增加条件！")
	}
	if len(list) > 1 {
		return errors.New("找到多条记录，请增加条件！")
	}
	return repository.OptionTradeRepo.Delete(&list[0])
}
