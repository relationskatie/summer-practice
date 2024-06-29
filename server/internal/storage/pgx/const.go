package pgx

const (
	queryMigrate = `CREATE TABLE IF NOT EXISTS vacancies
	(
		"id" UUID NOT NULL UNIQUE,
		"name" VARCHAR NOT NULL,
		"salary" VARCHAR NOT NULL,
		"area" VARCHAR NOT NULL,
		"url" VARCHAR NOT NULL,
		"employment" VARCHAR NOT NULL,
		"experience" VARCHAR NOT NULL,
		PRIMARY KEY("id")
	);`
	queryGetByID   = `SELECT * FROM vacancies where id = $1`
	queryGetAll    = `SELECT * FROM vacancies`
	queryDeleteAll = `DELETE FROM vacancies`
	queryAppend    = `INSERT INTO vacancies (id, name, salary, area, url, employment, experience) VALUES ($1, $2, $3, $4, $5, $6, $7)`
)
