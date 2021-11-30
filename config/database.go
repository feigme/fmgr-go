package config

type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns        int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode             string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	EnableFileLogWriter bool   `mapstructure:"enable-file-log-writer" json:"enable-file-log-writer" yaml:"enable-file-log-writer"`
	LogFilename         string `mapstructure:"log-filename" json:"log-filename" yaml:"log-filename"`
}
