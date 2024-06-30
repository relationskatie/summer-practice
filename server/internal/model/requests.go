package model

import uuid "github.com/vgarvardt/pgx-google-uuid/v5"

type FormRequest struct {
	ID         uuid.UUID `json:"id"`
	Text       string    `json"name"`
	Salary     string    `json:"salary"`
	Area       string    `json:"area`
	URL        string    `json:"url"`
	Employment string    `json:"employment"`
	Experience string    `json:"experience"`
}
