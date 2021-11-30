package config

type ApiServer struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	AppUrl  string `mapstructure:"app-url" json:"app-url" yaml:"app-url"`
}
