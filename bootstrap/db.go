package bootstrap

import (
	"os"

	"github.com/feigme/fmgr-go/global"
	"go.uber.org/zap"
)

// 数据库表初始化
func InitMySqlTables(dst ...interface{}) {

	err := global.App.DB.AutoMigrate(dst...)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
}
