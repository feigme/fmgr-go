package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/global"
)

type OptionTradeRepository struct{}

var OptionTradeRepo = new(OptionTradeRepository)

func (o *OptionTradeRepository) Save(trade *models.OptionTrade) error {
	return global.App.DB.Save(trade).Error
}

func (o *OptionTradeRepository) List(query *query.OptionTradeQuery) (list []models.OptionTrade) {
	tx := global.App.DB
	if query.Code != "" {
		code := strings.ToUpper(query.Code)
		tx = tx.Where(fmt.Sprintf(" code like '%%%s%%' ", code))
	}

	if len(query.StatusList) > 0 {
		tx = tx.Where(" status in (?)", query.StatusList)
	}

	tx.Find(&list)
	return list
}

func (o *OptionTradeRepository) Get(code string) (*models.OptionTrade, error) {
	var list []models.OptionTrade
	global.App.DB.Where(fmt.Sprintf(" code = '%s' ", code)).Find(&list)
	if len(list) == 0 {
		return nil, errors.New("期权code不存在！")
	}
	return &list[0], nil
}

func (o *OptionTradeRepository) Update(trade *models.OptionTrade) error {
	return global.App.DB.Updates(&trade).Error
}
