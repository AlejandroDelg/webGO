package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10

const maxIdleDBConn = 5

const maxDBLifeTime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	db, err := newDatabase(dsn)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = db
	err = testDB(dbConn)

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testDB(db *DB) error {
	err := db.SQL.Ping()
	if err != nil {
		return err
	}
	return nil
}

func newDatabase(dsn string) (*sql.DB, error) {
	db, error := sql.Open("pgx", dsn)

	if error != nil {
		return nil, error
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
