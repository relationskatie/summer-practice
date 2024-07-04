package storage

import (
	"context"

	"github.com/relationskatie/summer-practice/server/internal/model"
)

type VacanciesStorage interface {
	AppendAll(ctx context.Context, vacancies []model.ClientDTO) error
	GetAll(ctx context.Context) ([]model.ClientDTO, error)
	DeleteAll(ctx context.Context) error
}

type Interface interface {
	Vacancies() VacanciesStorage
}
