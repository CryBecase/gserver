package config

import (
	"flag"

	"github.com/spf13/viper"

	"gserver/pkg/cache/redis"
	"gserver/pkg/databse/mysql"
)

var (
	Conf     = &Config{}
	confPath string
)

type Config struct {
	Server *Server       `mapstructure:"server"`
	MySQL  *mysql.Config `mapstructure:"mysql"`
	Redis  *redis.Config `mapstructure:"redis"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Addr string `mapstructure:"addr"`
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

func Init() error {
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(Conf)
}
