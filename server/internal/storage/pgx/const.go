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
	queryGetByID = `SELECT * FROM vacancies where id = $1`
)
