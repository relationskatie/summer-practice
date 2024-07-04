package model

type ClientDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary struct {
		From int `json:"from"`
	} `json:"salary"`
	Area struct {
		Name string `json:"name"`
	} `json:"area"`
	URl string `json:"alternate_url"`
}

type ClientResponse struct {
	Items []ClientDTO `json:"items"`
}
