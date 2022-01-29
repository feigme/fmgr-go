package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/repository"
)

type OptionTradeService struct{}

var OptionTradeSvc = new(OptionTradeService)

func (svc *OptionTradeService) Save(trade *models.OptionTrade) error {
	return repository.NewOptionTradeRepo(context.Background()).Save(trade)
}

func (svc *OptionTradeService) List(query *query.OptionTradeQuery) (list []models.OptionTrade) {
	return repository.NewOptionTradeRepo(context.Background()).List(query)
}

func (svc *OptionTradeService) Get(code string) (*models.OptionTrade, error) {
	return repository.NewOptionTradeRepo(context.Background()).Get(code)
}

func (svc *OptionTradeService) GetById(id uint) (*models.OptionTrade, error) {
	return repository.NewOptionTradeRepo(context.Background()).GetById(id)
}

func (svc *OptionTradeService) Update(trade *models.OptionTrade) error {
	return repository.NewOptionTradeRepo(context.Background()).Update(trade)
}

func (svc *OptionTradeService) Delete(id uint) error {
	return repository.NewOptionTradeRepo(context.Background()).Delete(id)
}

func (svc *OptionTradeService) RollPut(id int, closePrice, exerciseDate, sellPrice string) error {
	trade, err := repository.NewOptionTradeRepo(context.Background()).GetById(uint(id))
	if err != nil {
		return errors.New(fmt.Sprintf("数据不存在，id=%d", id))
	}

	return repository.Transaction(context.Background(), func(txctx context.Context) error {
		trade.Close(closePrice)
		repository.NewOptionTradeRepo(context.Background()).Update(trade)

		code := strings.ReplaceAll(trade.Code, trade.ExerciseDate, exerciseDate)

		rollOptionTrade, err := models.NewOptionTrade(code, enum.Option_Position_Seller, sellPrice)
		if err != nil {
			return err
		}
		return repository.NewOptionTradeRepo(context.Background()).Save(rollOptionTrade)
	})
}
