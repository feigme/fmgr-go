package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/opst"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/repository"
)

type OpstService struct{}

var OpstSvc = new(OpstService)

func (svc *OpstService) Apply(opst *opst.Opst) error {
	return repository.Transaction(context.Background(), func(txctx context.Context) error {
		// 1. 策略名称检查
		stEnum, err := enum.GetOptionStEnumByName(opst.Spec.OptionStrategy)
		if err != nil {
			return errors.New(fmt.Sprintf("当前策略不支持, %s", opst.Spec.OptionStrategy))
		}

		// 2. 创建策略id
		st := models.NewOptionStrategy(stEnum)
		st.Name = opst.MetaData.Name
		st.Namespace = opst.MetaData.Namespace

		// todo 支持labels

		repository.NewOptionStrategyRepo(context.Background()).Save(st)

		// 3. 检查是否要正股

		// 4. 保存期权信息
		for _, op := range opst.Spec.Options {
			position, err := enum.GetOptionPositionEnumByName(op.Option.Position)
			if err != nil {
				return err
			}
			trade, err := models.NewOptionTrade(opst.Spec.Market, op.Option.Code, position, op.Option.Price)
			if err != nil {
				return err
			}

			// 跟策略关联
			trade.Sid = st.Id
			repository.NewOptionTradeRepo(context.Background()).Save(trade)
		}

		return nil
	})
}

func (svc *OpstService) List(query *query.OpstQuery) (list []models.OptionStrategy) {
	return repository.NewOptionStrategyRepo(context.Background()).List(query)
}
