package client

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Client struct {
	client *resty.Client
	log    *zap.Logger
}

func NewClient(log *zap.Logger) (*Client, error) {
	log.Info("Initialize Client")
	client := &Client{
		client: resty.New(),
		log:    log,
	}
	return client, nil
}
