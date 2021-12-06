package db

import (
	"os"
	"strconv"

	"github.com/feigme/fmgr-go/pkg/config"
	"github.com/feigme/fmgr-go/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func init() {
	// 根据驱动配置进行初始化
	switch config.Config.Database.Driver {
	case "mysql":
		initMySqlGorm()
	default:
		initMySqlGorm()
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm() *gorm.DB {
	dbConfig := config.Config.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,                                       // 禁用自动创建外键约束
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true}, // 禁用表名加s
	}); err != nil {
		log.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)

		DB = db
		log.Log.Info("db init success!")
		return db
	}
}

// 数据库表初始化
func InitMySqlTables(dst ...interface{}) {

	err := DB.AutoMigrate(dst...)
	if err != nil {
		log.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}

	// 程序关闭前，释放数据库连接
	defer func() {
		if DB != nil {
			db, _ := DB.DB()
			db.Close()
		}
	}()
}
