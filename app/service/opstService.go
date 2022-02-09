package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dianjiu/gokit/time"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/opst"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/repository"
	"github.com/feigme/fmgr-go/pkg/filehelper"
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
			trade, err := models.NewOptionTrade(opst.Spec.Market, op.Option.Code, position, op.Option.OpenPrice)
			if err != nil {
				return err
			}

			// 跟策略关联
			trade.Sid = st.Id

			if op.Option.OpenTime != "" {
				ot, err := time.ParseDateTime(op.Option.OpenTime)
				if err != nil {
					return err
				}
				trade.OpenTime = ot
			}

			if op.Option.CloseTime != "" {
				ct, err := time.ParseDateTime(op.Option.CloseTime)
				if err != nil {
					return err
				}
				trade.CloseTime = ct
			}

			if op.Option.ClosePrice == "0" {
				// 失效
				trade.Invalid()
			} else if op.Option.ClosePrice != "" {
				// 平仓
				trade.Close(op.Option.ClosePrice)
			}

			repository.NewOptionTradeRepo(context.Background()).Save(trade)
		}

		return nil
	})
}

func (svc *OpstService) List(query *query.OpstQuery) (list []models.OptionStrategy) {
	return repository.NewOptionStrategyRepo(context.Background()).List(query)
}

func (svc *OpstService) Import(path string) error {
	repo := repository.NewTradeHistoryRepo(context.Background())
	err := filehelper.ReadCsv(path, func(rows [][]string) error {
		n := 0
		for _, row := range rows {
			n = n + 1
			if n == 1 {
				continue
			}

			history := models.NewTradeHistory()
			history.Code = row[0]
			history.Name = row[1]
			history.Direction = row[2]
			history.Count = strings.ReplaceAll(row[3], ",", "")
			history.Price = strings.ReplaceAll(row[4], ",", "")
			history.Amount = strings.ReplaceAll(row[5], ",", "")
			tradeTime, _ := time.ParseDateTime(row[7])
			history.TradeTime = tradeTime

			if !repo.IsExisted(history) {
				repo.Save(history)
			}
		}
		return nil
	})
	return err
}
