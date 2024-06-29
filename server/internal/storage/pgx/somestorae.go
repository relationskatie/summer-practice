package pgx

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
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
	_, err := store.pool.Exec(context.Background(), queryMigrate)
	if err != nil {
		return err
	}
	return nil
}

func (store *vacanciesStorage) AppendAll(ctx context.Context, vacancies []model.ClientDTO) error {
	batch := &pgx.Batch{}
	for _, vacancy := range vacancies {
		batch.Queue(queryAppend, vacancy.ID, vacancy.Area, vacancy.Employment)
	}

	br := store.pool.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < len(vacancies); i++ {
		_, err := br.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgerrcode.UniqueViolation == pgErr.Code {
				return fmt.Errorf("unique violation error: %w", err)
			}
			return fmt.Errorf("error while inserting vacancies: %w", err)
		}
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
		err = row.Scan(&temp.ID)
		if err != nil {
			return nil, fmt.Errorf("error while scan vacancies: %w", err)
		}
		res = append(res, temp)
	}
	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: w%", err)
	}
	return res, nil
}

func (store *vacanciesStorage) GetVacancyById(ctx context.Context, id uuid.UUID) (*model.ClientDTO, error) {
	var vacancy model.ClientDTO
	err := store.pool.QueryRow(ctx, queryGetByID, id).Scan(&vacancy.ID, &vacancy.Name, &vacancy.Salary, &vacancy.Area.Name)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (store *vacanciesStorage) DeleteAll(ctx context.Context) error {
	_, err := store.pool.Exec(ctx, queryDeleteAll)
	if err != nil {
		return fmt.Errorf("error while deleting all vacancies: %w", err)
	}
	return nil
}
