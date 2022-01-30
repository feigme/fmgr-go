package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"gorm.io/gorm"
)

type OptionStrategyRepository struct {
	db *gorm.DB
}

func NewOptionStrategyRepo(ctx context.Context) *OptionStrategyRepository {
	return &OptionStrategyRepository{
		db: GetDB(ctx),
	}
}

func (repo *OptionStrategyRepository) Save(st *models.OptionStrategy) error {
	return repo.db.Save(st).Error
}

func (repo *OptionStrategyRepository) List(query *query.OpstQuery) (list []models.OptionStrategy) {
	tx := repo.db
	if query.Code != "" {
		code := strings.ToUpper(query.Code)
		tx = tx.Where(fmt.Sprintf(" code like '%%%s%%' ", code))
	}

	if len(query.StatusList) > 0 {
		tx = tx.Where(" status in (?)", query.StatusList)
	}

	tx.Offset(query.GetOffset()).Limit(query.PageSize).Order(" id desc").Find(&list)
	return list
}
