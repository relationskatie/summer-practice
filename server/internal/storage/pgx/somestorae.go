package pgx

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/summer-practice/server/internal/model"
	"go.uber.org/zap"
)

type vacanciesStorage struct {
	pool  *pgxpool.Pool
	log   *zap.Logger
	pgErr *pgconn.PgError
}

func newVacanciesStorage(pool *pgxpool.Pool, log *zap.Logger, pgErr *pgconn.PgError) (*vacanciesStorage, error) {
	store := &vacanciesStorage{
		pool:  pool,
		log:   log,
		pgErr: pgErr,
	}
	if err := store.migrate(); err != nil {
		return nil, err
	}
	return store, nil
}

func (store *vacanciesStorage) migrate() error {
	_, err := store.pool.Exec(context.Background(), quertMigrate)
	if err != nil {
		return err
	}
	return nil
}

func (store *vacanciesStorage) Append(ctx context.Context, vacancy *model.ClientDTO) error {
	_, err := store.pool.Exec(ctx, queryAppend)
	if err != nil {
		if errors.As(err, &store.pgErr) && pgerrcode.UniqueViolation == store.pgErr.Code {
			return err
		}
		return err
	}
	return nil
}

func (store *vacanciesStorage) GetAll(ctx context.Context) ([]model.ClientDTO, error) {
	var res []model.ClientDTO

	row, err := store.pool.Query(ctx, queryGetAll)
	if err != nil {
		return nil, fmt.Errorf("error while gel all vacancies: %W", err)
	}
	defer row.Close()

	for row.Next() {
		var temp model.ClientDTO
		err = rows.Scan(&temp.ID)
		if err != nil {
			return nil, fmt.Errorf("error while scan vacancies: %w", err)
		}
		res = append(res, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: w%", err)
	}
	return res, nil
}

func (store *vacanciesStorage) GetVacancyById(ctx context.Context, id uuid.UUID) (*model.ClientDTO, error) {
	var vacancy model.ClientDTO
	err := store.pool.QueryRow(ctx, queryGetByID, id).Scan(&vacancy.Area)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

//func (store *Storage)DeleteAll(ctx context.Context)error{

//}
