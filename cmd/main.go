package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime)

	serverLogger := server.NewServer(logger)

	err := serverLogger.Start()

	if err != nil {
		logger.Fatalf("Fatal server run: %v", err)
	}
}
