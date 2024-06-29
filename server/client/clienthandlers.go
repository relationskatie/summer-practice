package client

import (
	"encoding/json"
	"fmt"

	"github.com/relationskatie/summer-practice/server/internal/model"
	tokenapi "github.com/relationskatie/summer-practice/server/token"
	"go.uber.org/zap"
)

func (client *Client) handleClientDo() ([]model.ClientDTO, error) {
	//var req *model.ClientRequest
	var Vacancies model.ClientResponse
	tok := tokenapi.GetToken()
	resp, err := client.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", tok)).
		SetQueryParams(map[string]string{
			"text":             "разработчик",
			"area":             "1",
			"salary":           "1000000",
			"only_with_salary": "true",
			"per_page":         "20",
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
