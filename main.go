package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := LoadConfig()

	// Default port fallback
	if config.Port == "" {
		config.Port = "8080"
	}

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("gagal running server:", err)
	}
}
