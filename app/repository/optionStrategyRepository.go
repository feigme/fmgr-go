package repository

import (
	"context"

	"github.com/feigme/fmgr-go/app/models"
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
