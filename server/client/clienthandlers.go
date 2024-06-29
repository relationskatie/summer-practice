package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/relationskatie/summer-practice/server/internal/model"
	tokenapi "github.com/relationskatie/summer-practice/server/token"
	"go.uber.org/zap"
)

func (client *Client) HandleClientDo() error {
	var req *model.ClientRequest
	var Vacancies []model.ClientDTO
	tok := tokenapi.GetToken()
	resp, err := client.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", tok)).
		SetQueryParams(map[string]string{
			"text":             req.Text,
			"are":              req.Area,
			"salary":           req.Salary,
			"only_with_salary": strconv.FormatBool(req.OnlyWithSalary),
			"per_page":         strconv.Itoa(req.PerPage),
		}).
		Get("https://api.hh.ru/vacancies")
	if err != nil {
		client.log.Error("Failed to request hh.ru", zap.Error(err))
		return err
	}
	err = json.Unmarshal(resp.Body(), &Vacancies)
	if err != nil {
		client.log.Error("Failed to unmarshall result", zap.Error(err))
	}
	return err
}
