package config

import (
	"fmt"
	"os"

	"github.com/feigme/fmgr-go/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *config.Config

func init() {
	// 设置配置文件路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 1. 初始化 viper
	v := viper.New()
	// 2. 设置文件名称
	v.SetConfigFile(config)
	// 3. 配置类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("read config failed: %s \n", err)
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
}
