package test

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/feigme/fmgr-go/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Mock sqlmock.Sqlmock
)

func init() {
	var err error
	var conn *sql.DB
	conn, Mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("init sql mock failed, err: %v", err)
	}

	global.App.DB, err = gorm.Open(mysql.New(mysql.Config{SkipInitializeWithVersion: true, Conn: conn}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 禁用表名加s
	})

	if err != nil {
		log.Fatalf("init DB with sqlmock failed, err: %v", err)
	}
	global.App.Log.Info("start to mock db")
}
