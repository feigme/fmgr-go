package opst

import (
	"errors"

	"github.com/feigme/fmgr-go/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Opst struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   metaData `yaml:"metaData"`
	Spec       spec     `yaml:"spec"`
}

type metaData struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Labels    map[string]string `yaml:"labels"`
}

type spec struct {
	Market         string    `yaml:"market"`
	OptionStrategy string    `yaml:"optionStrategy"`
	Stock          stock     `yaml:"stock"`
	Options        []options `yaml:"options"`
}

type stock struct {
	Code  string `yaml:"code"`
	Cost  string `yaml:"cost"`
	Count uint   `yaml:"count"`
}

type options struct {
	Option option `yaml:"option"`
}

type option struct {
	Code         string `yaml:"code"`
	Position     string `yaml:"position"`
	Price        string `yaml:"price"`
	ClosePrice   string `yaml:"closePrice"`
	Count        uint   `yaml:"count"`
	ContractSize uint   `yaml:"contractSize"`
}

func ParseYaml(path string) (*Opst, error) {
	global.App.Log.Info("解析文件", zap.String("file", path))
	// 1. 初始化 viper
	v := viper.New()
	// 2. 设置文件名称
	v.SetConfigFile(path)
	// 3. 配置类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		global.App.Log.Error("读取文件失败", zap.Any("err", err))
		return nil, errors.New("读取文件失败")
	}
	// 4. 将配置赋值给全局变量
	var opst Opst
	if err := v.Unmarshal(&opst); err != nil {
		global.App.Log.Error("解析文件失败", zap.Any("err", err))
		return nil, errors.New("解析文件失败")
	}

	return &opst, nil
}
