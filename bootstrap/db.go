package bootstrap

import (
	"github.com/feigme/fmgr-go/pkg/db"
)

// 数据库表初始化
func InitMySqlTables(dst ...interface{}) {
	db.InitMySqlTables(dst...)
}
