package main

import (
	"context"
	"flag"
	"time"

	"go.uber.org/zap"

	"github.com/MihailPodgorny/go_concurrency_db/internal/config"
	"github.com/MihailPodgorny/go_concurrency_db/internal/tcpclient"
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
	msgSize := flag.String("max_message_size", "4KB", "Max message size for connection")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.NewClientConfig(*address, *timeout, *msgSize)
	if err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	client, err := tcpclient.NewTCPClient(cfg, logger)
	if err != nil {
		logger.Fatal("failed to connect with server", zap.Error(err))
	}

	err = client.Run(ctx)
	if err != nil {
		logger.Error("connection was closed", zap.Error(err))
		return err
	}

	return nil
}
