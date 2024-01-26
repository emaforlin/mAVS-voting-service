package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	Port uint16
}

type Db struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadConfig(); err != nil {
		panic(fmt.Errorf("error with config file: ", err))
	}

	return Config{
		App: App{
			Port: viper.GetUint16("app.server.port"),
		},
		Db: Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetUint16("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.passwd"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.tz"),
		},
	}

}
