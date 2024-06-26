package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/relationskatie/summer-practice/server/internal/model"
)

type Vacanciies interface {
	GetVacanciesByTunning(ctx context.Context, vacans *model.VacansDTO) ([]model.VacansDTO, error)
	GetVacancyById(ctx context.Context, id uuid.UUID) (*model.VacansDTO, error)
	GetAllFavourite(ctx context.Context) ([]model.VacansDTO, error)
}
