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

func (repo *OptionTradeRepository) Save(trade *models.OptionTrade) error {
	return global.App.DB.Save(trade).Error
}

func (repo *OptionTradeRepository) GetById(id uint) (trade *models.OptionTrade, err error) {
	err = global.App.DB.Where(" id = ? ", id).Find(&trade).Error
	return trade, err
}

func (repo *OptionTradeRepository) List(query *query.OptionTradeQuery) (list []models.OptionTrade) {
	tx := global.App.DB
	if query.Code != "" {
		code := strings.ToUpper(query.Code)
		tx = tx.Where(fmt.Sprintf(" code like '%%%s%%' ", code))
	}

	if len(query.StatusList) > 0 {
		tx = tx.Where(" status in (?)", query.StatusList)
	}

	if query.StartExerciseDate != "" {
		tx = tx.Where(" exercise_date >= ? ", query.StartExerciseDate)
	}
	if query.EndExerciseDate != "" {
		tx = tx.Where(" exercise_date <= ? ", query.EndExerciseDate)
	}

	tx.Offset(query.GetOffset()).Limit(query.PageSize).Order(" code asc, exercise_date asc").Find(&list)
	return list
}

func (repo *OptionTradeRepository) Get(code string) (*models.OptionTrade, error) {
	var list []models.OptionTrade
	global.App.DB.Where(fmt.Sprintf(" code = '%s' ", code)).Find(&list)
	if len(list) == 0 {
		return nil, errors.New("期权code不存在！")
	}
	return &list[0], nil
}

func (repo *OptionTradeRepository) Update(trade *models.OptionTrade) error {
	return global.App.DB.Updates(&trade).Error
}

func (repo *OptionTradeRepository) Delete(id uint) error {
	trade := &models.OptionTrade{}
	trade.Id = id
	return global.App.DB.Delete(trade).Error
}
