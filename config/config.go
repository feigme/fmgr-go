package config

type Config struct {
	ApiServer ApiServer `mapstructure:"api-server" json:"api-server" yaml:"api-server"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
	Database  Database  `mapstructure:"database" json:"database" yaml:"database"`
}
