package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() Config {
	// Load .env if exists
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// Read from environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return Config{
		Port: viper.GetString("PORT"),
	}
}
