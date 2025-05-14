package driver

import (
	"database/sql"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConns = 10
const maxIdleDbConns = 5
const maxLifetimeDbConns = 5 * time.Minute

func Connect(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxIdleConns(maxIdleDbConns)
	d.SetMaxOpenConns(maxOpenDbConns)
	d.SetConnMaxLifetime(maxLifetimeDbConns)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
