package model

import "github.com/google/uuid"

type ClientResponse struct {
	ID         uuid.UUID
	Text       string
	Area       string
	Salary     string
	Experience string
	URL        string
}
