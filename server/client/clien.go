package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/relationskatie/summer-practice/server/internal/model"
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

func GetDataFromClient(text string, salary string, area string) ([]model.ClientDTO, error) {
	var (
		log    *zap.Logger
		client *Client
		err    error
	)
	log, err = zap.NewProduction()
	defer log.Sync()
	if err != nil {
		log.Fatal("Failed to initilize logger", zap.Error(err))
		return nil, err
	}
	client, err = NewClient(log)
	if err != nil {
		log.Fatal("Failed to initilize client", zap.Error(err))
		return nil, err
	}
	vacancies, err := client.handleClientDo(text, salary, area)
	if err != nil {
		log.Error("Error handling client request", zap.Error(err))
		return nil, err
	}
	return vacancies, nil
}
