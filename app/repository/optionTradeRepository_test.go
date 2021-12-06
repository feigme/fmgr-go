package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/global"
	"github.com/feigme/fmgr-go/pkg/test"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var mock sqlmock.Sqlmock

func init() {
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("init sql mock failed, err: %v", err)
	}

	global.App.DB, err = gorm.Open(mysql.New(mysql.Config{SkipInitializeWithVersion: true, Conn: db}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 禁用表名加s
	})

	if err != nil {
		log.Fatalf("init DB with sqlmock failed, err: %v", err)
	}
}

func TestSave(t *testing.T) {
	// mock
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `option_trade`").WithArgs(test.AnyTime{}, test.AnyTime{}, "NIO211126P40000", "short", 40000, 0, "P", "NIO", "211126", 100, "", "1.52").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	option, err := models.NewOption("NIO211126P40000")
	require.NoError(t, err)

	trade, err := models.NewOptionTrade(option, enum.SHORT, "1.52")
	require.NoError(t, err)

	err = OptionTradeRepo.Save(trade)
	require.NoError(t, err)

	require.NotNil(t, trade.Id)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
