package basic

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"examine/sqlx-examine/config"
)

type Location struct {
	Id   int    `db:"id"`
	Addr string `db:"addr"`
	Name string `db:"name"`
}

type Place struct {
	Id         int    `db:"id"`
	LocationId int    `db:"location_id"`
	Name       string `db:"name"`
}

var (
	schema = `
CREATE SCHEMA IF NOT EXISTS sqlx_test;

SET SEARCH_PATH TO sqlx_test;

CREATE TABLE IF NOT EXISTS location (
  id integer PRIMARY KEY,
  addr text,
  name text
 );

CREATE TABLE IF NOT EXISTS place (
  id integer PRIMARY KEY,
  location_id integer REFERENCES location (id),
  name text
);
`
	cleanup = `
SET SEARCH_PATH TO sqlx_test;

DROP TABLE IF EXISTS place;

DROP TABLE IF EXISTS location;

DROP SCHEMA IF EXISTS sqlx_test;
`
	insertLocation = `INSERT INTO location (id, addr, name) VALUES ($1, $2, $3)`

	insertLocationNamed = `INSERT INTO location (id, addr, name) VALUES (:id, :addr, :name)`

	insertPlace = `INSERT INTO place (id, location_id, name) VALUES ($1, $2, $3)`

	insertPlaceNamed = `INSERT INTO place (id, location_id, name) VALUES (:id, :location_id, :name)`

	selectLocations = `SELECT id, addr, name from location`
)

func GetDB() (*sqlx.DB, error) {

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config info: %v", err)
	}

	db, err := sqlx.Connect("pgx", cfg.PgURL)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil

}
