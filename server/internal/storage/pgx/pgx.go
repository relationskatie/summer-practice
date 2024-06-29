package pgx

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/summer-practice/server/internal/storage"
	"go.uber.org/zap"
)

type Storage struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	vacancy *vacanciesStorage
	pgErr   *pgconn.PgError
}

func New(pool *pgxpool.Pool, log *zap.Logger, pgErr *pgconn.PgError) (*Storage, error) {
	vacancy, err := newVacanciesStorage(pool, log, pgErr)
	if err != nil {
		return nil, err
	}
	store := &Storage{
		pool:    pool,
		log:     log,
		vacancy: vacancy,
	}
	return store, nil
}

func (s *Storage) Vacancies() storage.VacanciesStorage {
	return s.vacancy
}
