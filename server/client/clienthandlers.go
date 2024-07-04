package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/relationskatie/summer-practice/server/internal/model"
	tokenapi "github.com/relationskatie/summer-practice/server/token"
	"go.uber.org/zap"
)

type Area struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Areas []Area `json:"areas"`
}

func (client *Client) handleClientDo(text string, salary string, area string, experience string) ([]model.ClientDTO, error) {
	var Vacancies model.ClientResponse
	url := "https://api.hh.ru/vacancies"
	tok := tokenapi.GetToken()
	resp, err := client.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", tok)).
		SetQueryParams(map[string]string{
			"text":             text,
			"salary":           salary,
			"area":             area,
			"experience":       experience,
			"only_with_salary": "true",
			"per_page":         "30",
		}).
		Get(url)
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

func (client *Client) handleGetIDArea(cityName string) (string, error) {
	url := "https://api.hh.ru/areas"
	tok := tokenapi.GetToken()
	resp, err := client.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", tok)).
		Get(url)
	if err != nil {
		client.log.Error("Failed to making request on area id", zap.Error(err))
		return "", nil
	}

	var areas []Area
	err = json.Unmarshal(resp.Body(), &areas)
	if err != nil {
		client.log.Error("Failed to unmarshalling response body", zap.Error(err))
		return "", nil
	}
	return findCityIDByName(areas, cityName), nil

}

func findCityIDByName(areas []Area, cityName string) string {
	for _, area := range areas {
		if strings.EqualFold(area.Name, cityName) {
			return area.ID
		}
		cityID := findCityIDByName(area.Areas, cityName)
		if cityID != "" {
			return cityID
		}
	}
	return ""
}
