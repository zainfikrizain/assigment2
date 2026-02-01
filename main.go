package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	type Config struct {
		Port string `mapstructure:"PORT"`
	}

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := Config{
		Port: viper.GetString("PORT"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
