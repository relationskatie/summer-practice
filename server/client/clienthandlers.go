package client

import (
	"encoding/json"
	"fmt"

	"github.com/relationskatie/summer-practice/server/internal/model"
	tokenapi "github.com/relationskatie/summer-practice/server/token"
	"go.uber.org/zap"
)

func (client *Client) handleClientDo(text string, salary string, area string, employment string, experience string) ([]model.ClientDTO, error) {
	var Vacancies model.ClientResponse
	tok := tokenapi.GetToken()
	resp, err := client.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", tok)).
		SetQueryParams(map[string]string{
			"text":             text,
			"salary":           salary,
			"area":             area,
			"employment":       employment,
			"experience":       experience,
			"only_with_salary": "true",
			"per_page":         "30",
		}).
		Get("https://api.hh.ru/vacancies")
	if err != nil {
		client.log.Error("Failed to request hh.ru", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &Vacancies)
	if err != nil {
		client.log.Error("Failed to unmarshall result", zap.Error(err))
		return nil, err
	}

	return Vacancies.Items, nil
}
