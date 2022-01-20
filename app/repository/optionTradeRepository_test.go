package repository

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/pkg/test"
)

func TestSaveOption(t *testing.T) {

	Convey("测试新建期权记录", t, func() {
		// mock
		test.Mock.ExpectBegin()
		test.Mock.ExpectExec("INSERT INTO `option_trade`").WithArgs(test.AnyTime{}, test.AnyTime{}, "NIO211126P40000", "NIO", "211126", "P", "40000", 100, 0, "short", "", "1.52", "", "", -1, "152.00", 1).WillReturnResult(sqlmock.NewResult(1, 1))
		test.Mock.ExpectCommit()

		trade, err := models.NewOptionTrade("NIO211126P40000", enum.SHORT, "1.52")
		So(err, ShouldBeNil)

		NewOptionTradeRepo(context.TODO()).Save(trade)
		So(err, ShouldBeNil)

		So(trade.Id, ShouldNotBeNil)

		err = test.Mock.ExpectationsWereMet()
		So(err, ShouldBeNil)
	})

}
