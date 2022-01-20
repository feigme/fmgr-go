package repository

import (
	"context"
	"reflect"

	"github.com/feigme/fmgr-go/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ctxTransactionKey struct{}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			global.App.Log.Error("unexpect context value type: %s", zap.Any("err", reflect.TypeOf(tx)))
			return nil
		}

		return tx
	}

	global.App.Log.Info("new db")
	return global.App.DB.WithContext(ctx)
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := GetDB(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}
