package main

import (
	"flag"
	"time"

	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error

	address := flag.String("address", "localhost:3223", "Database address")
	timeout := flag.Duration("timeout", time.Minute, "Idle timeout for connection")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	client, err := network.NewTCPClient(*address, maxMessageSize, *idleTimeout)
	if err != nil {
		logger.Fatal("failed to connect with server", zap.Error(err))
	}

	return nil
}
