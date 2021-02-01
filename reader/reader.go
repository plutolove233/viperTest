package reader

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	IpAddrees string
	Port      int
	UserName  string
	Password  string
	DBName    string
	Charset   string
}

func (_config *Config) Read() {
	fmt.Println("importing...")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("mysql.host", "127.0.0.1")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:%v", err)
	}
	_config.UserName = viper.GetString("mysql.username")
	_config.Port = viper.GetInt("mysql.ports")
	_config.Charset = viper.GetString("mysql.charset")
	_config.DBName = viper.GetString("mysql.database")
	_config.Password = viper.GetString("mysql.password")
	_config.IpAddrees = viper.GetString("mysql.host")
}
