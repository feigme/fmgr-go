package global

import (
	"github.com/feigme/fmgr-go/config"
	aconfig "github.com/feigme/fmgr-go/pkg/config"
	"github.com/feigme/fmgr-go/pkg/db"
	"github.com/feigme/fmgr-go/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      *config.Config
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = &Application{
	Config: aconfig.Config,
	Log:    log.Log,
	DB:     db.DB,
}
