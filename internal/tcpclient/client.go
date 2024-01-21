package tcpclient

import (
	"context"
	"fmt"
	"github.com/MihailPodgorny/go_concurrency_db/internal/config"
	"go.uber.org/zap"
	"net"
	"time"
)

type TCPClient struct {
	connection     net.Conn
	maxMessageSize uint64
	idleTimeout    time.Duration

	logger *zap.Logger
}

func NewTCPClient(cfg *config.ClientConfig, logger *zap.Logger) (*TCPClient, error) {
	connection, err := net.Dial("tcp", cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return &TCPClient{
		connection:     connection,
		maxMessageSize: cfg.MaxMsgSize,
		idleTimeout:    cfg.Timeout,
	}, nil
}

func (c *TCPClient) Run(ctx context.Context) error {
	return nil
}

func (c *TCPClient) Send(request []byte) ([]byte, error) {
	if err := c.connection.SetDeadline(time.Now().Add(c.idleTimeout)); err != nil {
		return nil, err
	}

	if _, err := c.connection.Write(request); err != nil {
		return nil, err
	}

	response := make([]byte, c.maxMessageSize)
	count, err := c.connection.Read(response)
	if err != nil {
		return nil, err
	}

	return response[:count], nil
}
