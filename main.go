package main

import (
	"fmt"
	"log"
)

func main() {
	config := LoadConfig()

	fmt.Printf("FINAL CONFIG: %+v\n", config)

	if config.DBDSN == "" {
		log.Fatal("DB_DSN is required")
	}

	fmt.Println("DB connected using:", config.DBDSN)
}
