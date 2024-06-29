package model

import "github.com/google/uuid"

type ClientDTO struct {
	ID         uuid.UUID
	Text       string `json:"text"`
	Area       string `json:"area"`
	Salary     string `json:"salary"`
	Experience string `json:"experience"`
	URL        string `json:"url"`
}

type VacansDTO struct {
}
