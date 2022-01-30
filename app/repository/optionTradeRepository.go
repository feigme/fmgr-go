package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"gorm.io/gorm"
)

type OptionTradeRepository struct {
	db *gorm.DB
}

func NewOptionTradeRepo(ctx context.Context) *OptionTradeRepository {
	return &OptionTradeRepository{
		db: GetDB(ctx),
	}
}

func (repo *OptionTradeRepository) Save(trade *models.OptionTrade) error {
	return repo.db.Save(trade).Error
}

func (repo *OptionTradeRepository) GetById(id uint) (trade *models.OptionTrade, err error) {
	err = repo.db.Where(" id = ? ", id).Find(&trade).Error
	return trade, err
}

func (repo *OptionTradeRepository) List(query *query.OptionQuery) (list []models.OptionTrade) {
	tx := repo.db
	if query.Code != "" {
		code := strings.ToUpper(query.Code)
		tx = tx.Where(fmt.Sprintf(" code like '%%%s%%' ", code))
	}

	if len(query.StatusList) > 0 {
		tx = tx.Where(" status in (?)", query.StatusList)
	}

	if query.Position != "" {
		tx = tx.Where(" position = ? ", query.Position)
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
	repo.db.Where(fmt.Sprintf(" code = '%s' ", code)).Find(&list)
	if len(list) == 0 {
		return nil, errors.New("期权code不存在! ")
	}
	return &list[0], nil
}

func (repo *OptionTradeRepository) Update(trade *models.OptionTrade) error {
	return repo.db.Updates(&trade).Error
}

func (repo *OptionTradeRepository) Delete(id uint) error {
	trade := &models.OptionTrade{}
	trade.Id = id
	return repo.db.Delete(trade).Error
}
