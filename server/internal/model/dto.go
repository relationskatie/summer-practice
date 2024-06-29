package model

type ClientDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Area struct {
		Name string `json:"name"`
	} `json:"area"`
	URl        string `json:"url"`
	Employment struct {
		Name string `json:"name"`
	}
	Experience struct {
		Name string `json:"name"`
	}
}

type ClientResponse struct {
	Items   []ClientDTO `json:"items"`
	Found   int         `json:"found"`
	Pages   int         `json:"pages"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
}
