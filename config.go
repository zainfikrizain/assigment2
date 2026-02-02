package main

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string
	DBDSN string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return Config{
		Port:  viper.GetString("PORT"),
		DBDSN: viper.GetString("DB_DSN"),
	}
}
