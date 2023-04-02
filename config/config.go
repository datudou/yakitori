package config

type MySQLConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type Server struct {
	ShutdownTimeout int    `mapstructure:"shutdown_timeout" json:"shutdown_timeout"`
	Addr            string `mapstructure:"addr" json:"addr"`
	HTTPTimeout     int    `mapstructure:"http_timeout" json:"http_timeout"`
}

type ServerConfig struct {
	MySQLInfo MySQLConfig `mapstructure:"mysql" json:"mysql"`
	Server    Server      `mapstructure:"server" json:"server"`
}

var ServerConf ServerConfig
