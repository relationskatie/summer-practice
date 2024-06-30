DROP TABLE IF EXISTS vacancies;
CREATE TABLE vacancies (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR,
    "salary" INT,
    "area" VARCHAR NOT NULL,
    "url" VARCHAR
);
