package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/pkg/test"
	"github.com/smartystreets/goconvey/convey"
)

func TestSaveStrategy(t *testing.T) {
	convey.Convey("测试保存策略", t, func() {
		// mock
		test.Mock.ExpectBegin()
		test.Mock.ExpectExec("INSERT INTO `option_strategy`").WithArgs(test.AnyTime{}, test.AnyTime{}, "xxx", "default", "", "covered call").WillReturnResult(sqlmock.NewResult(1, 1))
		test.Mock.ExpectCommit()

		st := &models.OptionStrategy{
			Name:      "xxx",
			Namespace: "default",
			Code:      enum.OST_Covered_Call.Name(),
		}
		err := NewOptionStrategyRepo(context.Background()).Save(st)
		convey.So(err, convey.ShouldBeNil)
	})
}
