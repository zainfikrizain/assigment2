package database

import (
	"log"
	"tugas02/database"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure: "PORT`
	DBConn string `mapstructure: "DB_CONN"`
}

func main() {
	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
}
