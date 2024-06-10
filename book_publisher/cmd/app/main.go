package main

import (
	"book_publisher/internal/config"
	"book_publisher/internal/server"
)

func main() {
	logger := config.NewZapLogger()
	defer logger.Sync()

	config := config.NewConfig(logger)
	config.LoadConfig()

	server := server.NewServer(logger, config)
	server.Serve()
}
