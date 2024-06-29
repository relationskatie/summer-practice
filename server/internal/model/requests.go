package model

type ClientRequest struct {
	Text           string
	Area           string
	Salary         string
	OnlyWithSalary bool
	PerPage        int
	Experience     string
}
