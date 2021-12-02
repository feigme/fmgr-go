package global

import (
	"github.com/feigme/fmgr-go/config"
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

var App = new(Application)
