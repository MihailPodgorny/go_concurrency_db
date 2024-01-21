package config

import (
	"github.com/MihailPodgorny/go_concurrency_db/internal/helpers"
	"time"
)

type ClientConfig struct {
	Address    string
	Timeout    time.Duration
	MaxMsgSize uint64
}

func NewClientConfig(addr string, t time.Duration, maxMsgSize string) (*ClientConfig, error) {

	mms, err := helpers.ConvertFromString(maxMsgSize)
	if err != nil {
		return nil, err
	}

	return &ClientConfig{
		Address:    addr,
		Timeout:    t,
		MaxMsgSize: mms,
	}, nil
}
